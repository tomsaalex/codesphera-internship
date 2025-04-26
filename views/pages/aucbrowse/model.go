package aucbrowse

import (
	"context"
	"curs1_boilerplate/views/components/navbar"
	"io"
)

type ViewModel struct {
	navbar *navbar.Model
}

func MakeAuctionBrowsePage(navbar *navbar.Model) *ViewModel {
	return &ViewModel{
		navbar: navbar,
	}
}

func (vm *ViewModel) Render(ctx context.Context, w io.Writer) error {
	return BrowseAuctions(vm, vm.navbar).Render(ctx, w)
}
