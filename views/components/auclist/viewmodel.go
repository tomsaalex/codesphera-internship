package auclist

import (
	"context"
	"curs1_boilerplate/model"
	"curs1_boilerplate/views/components/auccard"
	"curs1_boilerplate/views/components/pagenav"
	"io"
)

type ViewModel struct {
	auctionCards []*auccard.Model
	pageNav      pagenav.ViewModel
}

func MakeStandardAuctionList(auctions []model.Auction, pageNav pagenav.ViewModel) *ViewModel {
	auctionCards := make([]*auccard.Model, len(auctions))

	for i, auction := range auctions {
		auctionCards[i] = auccard.MakeAuctionCard(auction)
	}

	return &ViewModel{
		auctionCards: auctionCards,
		pageNav:      pageNav,
	}
}

func (vm *ViewModel) Render(ctx context.Context, w io.Writer) error {
	return PaginatedAuctionList(vm).Render(ctx, w)
}
