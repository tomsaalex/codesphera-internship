package model

import "github.com/google/uuid"

type Bid struct {
	Id    uuid.UUID
	Value float32

	RefAuction *Auction
	Bidder     *User
}
