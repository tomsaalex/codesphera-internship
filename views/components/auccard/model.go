package auccard

import (
	"context"
	"curs1_boilerplate/model"
	"io"
)

type Model struct {
	auction model.Auction
}

func MakeAuctionCard(auction model.Auction) *Model {
	return &Model{
		auction: auction,
	}
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return AuctionCard(m).Render(ctx, w)
}
