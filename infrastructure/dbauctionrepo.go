package infrastructure

import (
	"context"
	"curs1_boilerplate/db"
	"curs1_boilerplate/model"
)

type DBAuctionRepo struct {
	queries *db.Queries
	mapper  EntityMapperDB
}

func (r *DBAuctionRepo) GetAllAuctionsByUser(ctx context.Context, seller model.User) ([]model.Auction, error) {
	dbAuctions, err := r.queries.GetAllAuctionsByUser(ctx, r.mapper.uuidToDBUuid(seller.Id))
	if err != nil {
		// TODO: Differentiate them based on error codes later.
		return nil, &RepositoryError{Message: "database couldn't retrieve auctions"}
	}

	auctions, err := r.mapper.DBAuctionsToAuction(dbAuctions, &seller)
	if err != nil {
		return nil, &RepositoryError{Message: "entity couldn't be mapped to DB model"}
	}

	return auctions, nil
}

func (r *DBAuctionRepo) Add(ctx context.Context, auction model.Auction) (*model.Auction, error) {
	addAuctionParams, err := r.mapper.AuctionToAddAuctionParams(auction)
	if err != nil {
		return nil, &RepositoryError{Message: "entity couldn't be mapped to DB model"}
	}
	dbAuction, err := r.queries.AddAuction(ctx, *addAuctionParams)
	if err != nil {
		// TODO: Differentiate them based on error codes later.
		return nil, &RepositoryError{Message: "database couldn't accept auction"}
	}

	modelAuction, err := r.mapper.DBAuctionToAuction(dbAuction, auction.Seller)
	if err != nil {
		return nil, &RepositoryError{Message: "couldn't return added auction"}
	}

	return modelAuction, nil
}
