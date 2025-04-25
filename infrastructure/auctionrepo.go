package infrastructure

import (
	"context"
	"curs1_boilerplate/db"
	"curs1_boilerplate/model"
)

type AuctionRepository interface {
	GetAllAuctionsByUser(ctx context.Context, seller model.User) ([]model.Auction, error)
	GetAuctionByName(ctx context.Context, productName string) (*model.Auction, error)
	Add(ctx context.Context, newAuction model.Auction) (*model.Auction, error)
}

func NewDBAuctionRepository(queries *db.Queries) AuctionRepository {
	return &DBAuctionRepo{
		queries: queries,
		mapper:  EntityMapperDB{},
	}
}
