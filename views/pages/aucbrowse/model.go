package aucbrowse

import (
	"context"
	"curs1_boilerplate/model"
	"curs1_boilerplate/views/components/auclist"
	"curs1_boilerplate/views/components/navbar"
	"io"
)

type ViewModel struct {
	auctionList auclist.ViewModel
	categories  []model.Category

	navbar *navbar.Model
}

func MakeAuctionBrowsePage(auctionList auclist.ViewModel, categories []model.Category, navbar *navbar.Model) *ViewModel {
	return &ViewModel{
		auctionList: auctionList,
		navbar:      navbar,
		categories:  categories,
	}
}

func (vm *ViewModel) Render(ctx context.Context, w io.Writer) error {
	return BrowseAuctions(vm, vm.navbar).Render(ctx, w)
}
