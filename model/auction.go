package model

import "github.com/google/uuid"

type AuctionMode int
type AuctionStatus int

const (
	AM_Manual    AuctionMode = 0
	AM_Price_Met AuctionMode = 1
)

const (
	AS_Ongoing  AuctionStatus = 0
	AS_Finished AuctionStatus = 1
)

type Auction struct {
	Id                 uuid.UUID
	ProductName        string
	ProductDescription string
	Status             AuctionStatus
	Mode               AuctionMode
	StartingPrice      *float32
	TargetPrice        *float32
	//Categories         []string
	//ImageLinks         []string

	Seller *User
}
