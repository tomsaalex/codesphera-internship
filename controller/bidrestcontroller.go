package controller

import (
	"curs1_boilerplate/service"
	"curs1_boilerplate/util"
)

type BidRestController struct {
	bidService service.BidService
	jwtHelper  util.JwtUtil
}

func NewBidRestController(bidService service.BidService, jwtHelper util.JwtUtil) *BidRestController {
	return &BidRestController{
		bidService: bidService,
		jwtHelper:  jwtHelper,
	}
}
