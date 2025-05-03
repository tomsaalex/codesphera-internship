package aucbrowse

import (
	"context"
	"curs1_boilerplate/model"
	custalerts "curs1_boilerplate/views/components/alert"
	"curs1_boilerplate/views/components/auclist"
	"curs1_boilerplate/views/components/navbar"
	"io"
)

type ViewModel struct {
	auctionList *auclist.ViewModel
	categories  []model.Category

	navbar *navbar.Model
	alert  *custalerts.ViewModel
}

func MakeAuctionBrowsePage(auctionList *auclist.ViewModel, categories []model.Category, navbar *navbar.Model, alert *custalerts.ViewModel) *ViewModel {
	return &ViewModel{
		auctionList: auctionList,
		navbar:      navbar,
		categories:  categories,
		alert:       alert,
	}
}

func (vm *ViewModel) Render(ctx context.Context, w io.Writer) error {
	return BrowseAuctions(vm, vm.navbar).Render(ctx, w)
}
