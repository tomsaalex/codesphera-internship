package service

import "curs1_boilerplate/infrastructure"

type BidService struct {
	bidRepo infrastructure.BidRepository
}

func NewBidService() *BidService {
	return &BidService{}
}
