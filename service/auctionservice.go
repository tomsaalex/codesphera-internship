package service

import "curs1_boilerplate/infrastructure"

type AuctionService struct {
	auctionRepo infrastructure.AuctionRepository
	dtoMapper   ServiceDTOMapper
}

func NewAuctionService(auctionRepo infrastructure.AuctionRepository, dtoMapper ServiceDTOMapper) *AuctionService {
	return &AuctionService{
		auctionRepo: auctionRepo,
		dtoMapper:   dtoMapper,
	}
}
