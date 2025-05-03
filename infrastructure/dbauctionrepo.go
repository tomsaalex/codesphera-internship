package infrastructure

import (
	"context"
	"curs1_boilerplate/db"
	"curs1_boilerplate/model"
	"curs1_boilerplate/sharederrors"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

type DBAuctionRepo struct {
	queries *db.Queries
	mapper  EntityMapperDB
}

func (r *DBAuctionRepo) GetAuctions(ctx context.Context, auctionFilter AuctionFilter) ([]model.Auction, int, error) {
	auctionSearchParams, err := r.mapper.AuctionFilterToGetAuctionParams(auctionFilter)

	if err != nil {
		return nil, 0, &RepositoryError{Message: "Failed to convert auction search parameters to DB struct"}
	}

	dbAuctions, err := r.queries.GetAuctions(ctx, auctionSearchParams)
	if err != nil {
		return nil, 0, &RepositoryError{Message: err.Error()}
	}

	auctions, totalMatchingAuctions, err := r.mapper.DBAuctionDetailsToAuctions(dbAuctions)
	if err != nil {
		return nil, 0, &RepositoryError{Message: "Couldn't map db auctions to domain auctions"}
	}

	return auctions, totalMatchingAuctions, nil
}

func (r *DBAuctionRepo) GetAllAuctionsByUser(ctx context.Context, seller model.User) ([]model.Auction, error) {
	dbAuctions, err := r.queries.GetAllAuctionsByUser(ctx, r.mapper.uuidToDBUuid(seller.Id))
	if err != nil {
		// TODO: Differentiate them based on error codes later.
		return nil, &RepositoryError{Message: "database couldn't retrieve auctions"}
	}

	auctions, _, err := r.mapper.DBAuctionDetailsToAuctions(dbAuctions)
	if err != nil {
		return nil, &RepositoryError{Message: "entity couldn't be mapped to DB model"}
	}

	return auctions, nil
}

func (r *DBAuctionRepo) GetAuctionByName(ctx context.Context, productName string) (*model.Auction, error) {
	dbAuction, err := r.queries.GetAuctionByName(ctx, productName)
	if err != nil {
		return nil, &RepositoryError{Message: "database couldn't retrieve auction"}
	}

	auction, err := r.mapper.DBAuctionDetailToAuction(dbAuction)

	if err != nil {
		return nil, &RepositoryError{Message: "entity couldn't be mapped to DB model"}
	}

	return auction, nil
}

func (r *DBAuctionRepo) Add(ctx context.Context, auction model.Auction) (*model.Auction, error) {
	addAuctionParams, err := r.mapper.AuctionToAddAuctionParams(auction)

	if err != nil {
		return nil, &EntityDBMappingError{Message: "entity couldn't be mapped to DB model"}
	}
	dbAuction, err := r.queries.AddAuction(ctx, *addAuctionParams)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				return nil, &sharederrors.DuplicateEntityError{Message: "another auction has that same name"}
			case "23503":
				return nil, &ForeignKeyViolationError{Message: "no user found that has the id mentioned as seller_id"}
			default:
				return nil, &RepositoryError{Message: "couldn't add auction to db"}
			}
		}
	}

	modelAuction, err := r.mapper.DBAuctionToAuction(dbAuction, auction.Category, auction.Seller)
	if err != nil {
		return nil, &RepositoryError{Message: "couldn't return added auction"}
	}

	return modelAuction, nil
}

func (r *DBAuctionRepo) GetCategories(ctx context.Context) ([]model.Category, error) {
	dbCategories, err := r.queries.GetCategories(ctx)
	if err != nil {
		return nil, &RepositoryError{Message: "Couldn't load categories"}
	}

	return r.mapper.DBCategoriesToCategories(dbCategories)
}
