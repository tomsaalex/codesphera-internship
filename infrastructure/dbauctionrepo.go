package infrastructure

import (
	"context"
	"curs1_boilerplate/db"
	"curs1_boilerplate/middleware"
	"curs1_boilerplate/model"
	"curs1_boilerplate/sharederrors"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBAuctionRepo struct {
	queries *db.Queries
	mapper  EntityMapperDB
}

func (r *DBAuctionRepo) GetAuctions(ctx context.Context, auctionFilter AuctionFilter) ([]model.Auction, int, error) {
	logger := middleware.LoggerFromContext(ctx).With(slog.String("Layer", "DBAuctionRepo"))
	auctionSearchParams, err := r.mapper.AuctionFilterToGetAuctionParams(auctionFilter)

	if err != nil {
		logger.Error("Failed to convert Auction to DB model")
		return nil, 0, &EntityDBMappingError{Message: "Failed to convert auction search parameters to DB struct"}
	}

	dbAuctions, err := r.queries.GetAuctions(ctx, auctionSearchParams)

	if err != nil {
		logger.Error("Failed to add Auction to DB")
		return nil, 0, &RepositoryError{Message: err.Error()}
	}

	auctions, totalMatchingAuctions, err := r.mapper.DBGetAuctionsRowsToAuctions(dbAuctions)
	if err != nil {
		logger.Warn("Failed to convert addition result to domain entity. Addition succeeded.")
		return nil, 0, &RepositoryError{Message: "Couldn't map db auctions to domain auctions"}
	}

	return auctions, totalMatchingAuctions, nil
}

func (r *DBAuctionRepo) GetAllAuctionsByUser(ctx context.Context, seller model.User) ([]model.Auction, error) {
	logger := middleware.LoggerFromContext(ctx).With(slog.String("Layer", "DBAuctionRepo"), slog.String("Auction Seller", seller.Email))

	dbAuctions, err := r.queries.GetAllAuctionsByUser(ctx, r.mapper.uuidToDBUuid(seller.Id))
	if err != nil {
		logger.Error("Failed to retrieve the Auctions belonging to a user")
		// TODO: Differentiate them based on error codes later.
		return nil, &RepositoryError{Message: "database couldn't retrieve auctions"}
	}

	auctions, err := r.mapper.DBAuctionDetailsToAuctions(dbAuctions)
	if err != nil {
		logger.Error("Auctions were retrieved successfully, but couldn't be mapped to domain entity")
		return nil, &RepositoryError{Message: "entity couldn't be mapped to DB model"}
	}

	return auctions, nil
}

func (r *DBAuctionRepo) GetAuctionByName(ctx context.Context, productName string) (*model.Auction, error) {
	logger := middleware.LoggerFromContext(ctx).With(slog.String("Layer", "DBAuctionRepo"), slog.String("Auction Name", productName))
	dbAuction, err := r.queries.GetAuctionByName(ctx, productName)
	if err != nil {
		logger.Error("Couldn't retrieve Auction by name")
		return nil, &RepositoryError{Message: "database couldn't retrieve auction"}
	}

	auction, err := r.mapper.DBAuctionDetailToAuction(dbAuction)

	if err != nil {
		logger.Error("Auction was retrieved successfully, but couldn't be mapped to the domain entity")
		return nil, &RepositoryError{Message: "Auction couldn't be mapped to domain model"}
	}

	return auction, nil
}

func (r *DBAuctionRepo) GetAuctionById(ctx context.Context, auctionId uuid.UUID) (*model.Auction, error) {
	logger := middleware.LoggerFromContext(ctx).With(slog.String("Layer", "DBAuctionRepo"), slog.String("Auction ID", auctionId.String()))

	dbUuid := r.mapper.uuidToDBUuid(auctionId)

	dbAuction, err := r.queries.GetAuctionById(ctx, dbUuid)

	if err != nil {
		logger.Error("No Auction was found with the given Id.")
		return nil, &RepositoryError{Message: "no Auction found with the given Id"}
	}

	auction, err := r.mapper.DBAuctionDetailToAuction(dbAuction)

	if err != nil {
		logger.Error("Auction couldn't be mapped to domain model.")
		return nil, &RepositoryError{Message: "Auction couldn't be mapped to domain model"}
	}

	return auction, nil
}

func (r *DBAuctionRepo) Add(ctx context.Context, auction model.Auction) (*model.Auction, error) {
	logger := middleware.LoggerFromContext(ctx).With(slog.String("Layer", "DBAuctionRepo")).With(slog.Any("Auction", auction))
	addAuctionParams, err := r.mapper.AuctionToAddAuctionParams(auction)

	if err != nil {
		logger.Error("Couldn't map auction to DB model")
		return nil, &EntityDBMappingError{Message: "entity couldn't be mapped to DB model"}
	}
	dbAuction, err := r.queries.AddAuction(ctx, *addAuctionParams)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				logger.Error("Auction is duplicated")
				return nil, &sharederrors.DuplicateEntityError{Message: "another auction has that same name"}
			case "23503":
				logger.Error("Auction seller couldn't be found")
				return nil, &ForeignKeyViolationError{Message: "no user found that has the id mentioned as seller_id"}
			default:
				logger.Error("Failed to add Auction to the DB for unknown reasons")
				return nil, &RepositoryError{Message: "couldn't add auction to db"}
			}
		}
	}

	modelAuction, err := r.mapper.DBAuctionToAuction(dbAuction, auction.Category, auction.Seller)
	if err != nil {
		logger.Error("Auction has been added, but the added Auction couldn't be mapped to a domain entity")
		return nil, &RepositoryError{Message: "couldn't return added auction"}
	}

	return modelAuction, nil
}

func (r *DBAuctionRepo) GetCategories(ctx context.Context) ([]model.Category, error) {
	logger := middleware.LoggerFromContext(ctx).With(slog.String("Layer", "DBAuctionRepo"))
	dbCategories, err := r.queries.GetCategories(ctx)
	if err != nil {
		logger.Error("Failed to retrieve Auction categories from the DB")
		return nil, &RepositoryError{Message: "Couldn't load categories"}
	}

	return r.mapper.DBCategoriesToCategories(dbCategories)
}
