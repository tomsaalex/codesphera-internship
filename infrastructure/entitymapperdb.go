package infrastructure

import (
	"curs1_boilerplate/db"
	"curs1_boilerplate/model"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// TODO: Consider splitting this into multiple mappers, each for one entity type.
type EntityMapperDB struct {
}

// User

func (d *EntityMapperDB) DBUserToUser(dbUser db.User) *model.User {
	user := model.User{
		Id:       dbUser.ID.Bytes,
		Email:    dbUser.Email,
		Fullname: dbUser.Fullname,
		PassHash: dbUser.PassHash,
		PassSalt: dbUser.PassSalt,
	}

	return &user
}

func (d *EntityMapperDB) UserToAddUserParams(user model.User) db.AddUserParams {
	return db.AddUserParams{
		Fullname: user.Fullname,
		Email:    user.Email,
		PassHash: user.PassHash,
		PassSalt: user.PassSalt,
	}
}

func (d *EntityMapperDB) UserToUpdateUserParams(user model.User) db.UpdateUserParams {
	return db.UpdateUserParams{
		Fullname: user.Fullname,
		Email:    user.Email,
		PassHash: user.PassHash,
		PassSalt: user.PassSalt,
	}
}

// Auction

func (d *EntityMapperDB) DBAuctionsToAuction(dbAuctions []db.Auction, seller *model.User) ([]model.Auction, error) {
	modelAuctions := make([]model.Auction, len(dbAuctions))
	for i, dbAuction := range dbAuctions {
		modelAuction, err := d.DBAuctionToAuction(dbAuction, seller)
		modelAuctions[i] = *modelAuction
		return nil, err
	}
	return modelAuctions, nil
}

func (d *EntityMapperDB) DBAuctionToAuction(dbAuction db.Auction, seller *model.User) (*model.Auction, error) {
	dbAuctionStatus, err := d.dbAuctionStatusToAuctionStatus(dbAuction.AucStatus)
	if err != nil {
		return nil, err
	}

	dbAuctionMode, err := d.dbAuctionModeToAuctionMode(dbAuction.AucMode)
	if err != nil {
		return nil, err
	}

	if !dbAuction.ID.Valid {
		return nil, fmt.Errorf("couldn't convert db auction to model Auction - invalid ID.")
	}
	auctionUUID, err := uuid.FromBytes(dbAuction.ID.Bytes[:])
	if err != nil {
		return nil, err
	}

	if !dbAuction.TargetPrice.Valid {
		return nil, fmt.Errorf("couldn't convert db auction to model Auction - invalid target price")
	}
	targetPrice := dbAuction.TargetPrice.Float32

	auction := &model.Auction{
		Id:                 auctionUUID,
		ProductName:        dbAuction.ProductName,
		ProductDescription: dbAuction.ProductDesc,
		Mode:               dbAuctionMode,
		Status:             dbAuctionStatus,
		StartingPrice:      dbAuction.StartingPrice,
		TargetPrice:        targetPrice,
		Seller:             seller,
	}

	return auction, nil
}

func (d *EntityMapperDB) uuidToDBUuid(modelUuid uuid.UUID) pgtype.UUID {
	// Perhaps it looks bizarre, but Google's uuid is just syntax sugar over a [16]byte, so it checks out.
	return pgtype.UUID{
		Bytes: modelUuid,
		Valid: true,
	}
}

func (d *EntityMapperDB) auctionModeToDBAuctionMode(aucMode model.AuctionMode) (db.AuctionMode, error) {
	switch aucMode {
	case model.AM_Manual:
		return db.AuctionModeManual, nil
	case model.AM_Price_Met:
		return db.AuctionModePriceMet, nil
	default:
		return db.AuctionModeManual, fmt.Errorf("couldn't convert AuctionMode enum to db variant")
	}
}

func (d *EntityMapperDB) dbAuctionModeToAuctionMode(aucMode db.AuctionMode) (model.AuctionMode, error) {
	switch aucMode {
	case db.AuctionModeManual:
		return model.AM_Manual, nil
	case db.AuctionModePriceMet:
		return model.AM_Price_Met, nil
	default:
		return model.AM_Manual, fmt.Errorf("couldn't convert db AuctionMode enum to model")
	}
}

func (d *EntityMapperDB) auctionStatusToDBAuctionStatus(aucStatus model.AuctionStatus) (db.AuctionStatus, error) {
	switch aucStatus {
	case model.AS_Ongoing:
		return db.AuctionStatusOngoing, nil
	case model.AS_Finished:
		return db.AuctionStatusFinished, nil
	default:
		return db.AuctionStatusOngoing, fmt.Errorf("couldn't convert AuctionStatus enum to db variant")
	}
}

func (d *EntityMapperDB) dbAuctionStatusToAuctionStatus(aucStatus db.AuctionStatus) (model.AuctionStatus, error) {
	switch aucStatus {
	case db.AuctionStatusOngoing:
		return model.AS_Ongoing, nil
	case db.AuctionStatusFinished:
		return model.AS_Finished, nil
	default:
		return model.AS_Ongoing, fmt.Errorf("couldn't convert db AuctionStatus enum to model variant")
	}
}

func (d *EntityMapperDB) AuctionToAddAuctionParams(auction model.Auction) (*db.AddAuctionParams, error) {
	dbTargetPrice := pgtype.Float4{}
	err := dbTargetPrice.Scan(auction.TargetPrice)
	if err != nil {
		return nil, err
	}

	dbAuctionStatus, err := d.auctionStatusToDBAuctionStatus(auction.Status)
	if err != nil {
		return nil, err
	}

	dbAuctionMode, err := d.auctionModeToDBAuctionMode(auction.Mode)
	if err != nil {
		return nil, err
	}

	pgSellerUUID := pgtype.UUID{}
	err = pgSellerUUID.Scan(auction.Seller.Id)

	if err != nil {
		return nil, err
	}

	return &db.AddAuctionParams{
		ProductName:   auction.ProductName,
		ProductDesc:   auction.ProductDescription,
		AucMode:       dbAuctionMode,
		AucStatus:     dbAuctionStatus,
		StartingPrice: auction.StartingPrice,
		TargetPrice:   dbTargetPrice,
		SellerID:      pgSellerUUID,
	}, nil
}
