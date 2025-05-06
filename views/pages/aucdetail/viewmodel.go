package aucdetail

import (
	"context"
	"curs1_boilerplate/middleware"
	"curs1_boilerplate/model"
	"curs1_boilerplate/views/components/navbar"
	"io"
)

type ViewModel struct {
	loggedIn bool
	auction  model.Auction
	navbar   *navbar.Model
}

func MakeAuctionDetailPage(ctx context.Context, auction *model.Auction, navbar *navbar.Model) *ViewModel {
	userEmail := middleware.GetUserEmailFromContext(ctx)

	loggedIn := userEmail != ""

	return &ViewModel{
		loggedIn: loggedIn,
		auction:  *auction,
		navbar:   navbar,
	}
}

func (vm *ViewModel) Render(ctx context.Context, w io.Writer) error {
	return AuctionDetailPage(vm, vm.navbar).Render(ctx, w)
}
