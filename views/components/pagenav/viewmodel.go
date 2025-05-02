package pagenav

import (
	"context"
	"fmt"
	"io"
	"strconv"
)

type pageLink struct {
	pageNum    string
	isSelected bool
}

func (l *pageLink) GetLink() string {
	return fmt.Sprintf("/whatever/%s", l.pageNum)
}

type ViewModel struct {
	pageCount    int
	selectedPage int
	pages        []pageLink
}

func generatePageLinks(pageCount, selectedPage int) []pageLink {
	var pagesToDisplay []pageLink
	if pageCount == 1 {
		pagesToDisplay = []pageLink{{pageNum: "1", isSelected: true}}
	} else if pageCount == 2 {
		pagesToDisplay = []pageLink{
			{pageNum: "1", isSelected: selectedPage == 1},
			{pageNum: "2", isSelected: selectedPage == 2},
		}
	} else {
		if selectedPage == 1 {
			pagesToDisplay = []pageLink{
				{pageNum: "1", isSelected: true},
				{pageNum: "2", isSelected: false},
				{pageNum: "3", isSelected: false},
			}
		} else if selectedPage == pageCount {
			pagesToDisplay = []pageLink{
				{pageNum: strconv.Itoa(selectedPage - 2), isSelected: false},
				{pageNum: strconv.Itoa(selectedPage - 1), isSelected: false},
				{pageNum: strconv.Itoa(selectedPage), isSelected: true},
			}
		} else {
			pagesToDisplay = []pageLink{
				{pageNum: strconv.Itoa(selectedPage - 1), isSelected: false},
				{pageNum: strconv.Itoa(selectedPage), isSelected: true},
				{pageNum: strconv.Itoa(selectedPage + 1), isSelected: false},
			}
		}
	}

	return pagesToDisplay
}

func MakePageNav(pageCount, selectedPage int) *ViewModel {
	pagesToDisplay := generatePageLinks(pageCount, selectedPage)

	return &ViewModel{
		pageCount:    pageCount,
		pages:        pagesToDisplay,
		selectedPage: selectedPage,
	}
}

func (vm *ViewModel) Render(ctx context.Context, w io.Writer) error {
	return PaginationNav(vm).Render(ctx, w)
}
