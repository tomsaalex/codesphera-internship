package aucbrowse

import (
	"context"
	"curs1_boilerplate/model"
	"curs1_boilerplate/views/components/auccard"
	"curs1_boilerplate/views/components/navbar"
	"io"
)

type ViewModel struct {
	auctionCards []*auccard.Model
	navbar       *navbar.Model
}

func MakeAuctionBrowsePage(auctions []model.Auction, navbar *navbar.Model) *ViewModel {
	auctionCards := make([]*auccard.Model, len(auctions))

	for i, auction := range auctions {
		auctionCards[i] = auccard.MakeAuctionCard(auction)
	}

	return &ViewModel{
		auctionCards: auctionCards,
		navbar:       navbar,
	}
}

func (vm *ViewModel) Render(ctx context.Context, w io.Writer) error {
	return BrowseAuctions(vm, vm.navbar).Render(ctx, w)
}
