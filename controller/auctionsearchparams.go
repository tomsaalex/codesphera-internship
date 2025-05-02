package controller

import (
	"curs1_boilerplate/infrastructure"
)

type AuctionSearchParams struct {
	ProductQuery string `json:"productQuery"`
	CategoryName string `json:"categoryName"`

	Reverse string `json:"reverse"`
	OrderBy string `json:"orderBy"`

	SkippedPages int `json:"skippedPages"`
	PageSize     int `json:"pageSize"`
}

func (asp *AuctionSearchParams) ToServiceStruct() *infrastructure.AuctionFilter {
	skippedPages := asp.SkippedPages
	if skippedPages <= 0 {
		skippedPages = 0
	}

	pageSize := asp.PageSize
	if pageSize <= 0 {
		pageSize = 5
	}

	// TODO: Redo the whole bool/string thing
	var reverse bool
	if asp.Reverse == "true" {
		reverse = true
	} else {
		reverse = false
	}

	var orderBy infrastructure.AuctionOrderParameter
	switch asp.OrderBy {
	case "Product Name":
		orderBy = infrastructure.AuctionOrderName
	case "Created At":
		orderBy = infrastructure.AuctionOrderCreatedAt
	default:
		orderBy = infrastructure.AuctionOrderCreatedAt
	}

	return &infrastructure.AuctionFilter{
		ProductName:  asp.ProductQuery,
		ProductDesc:  asp.ProductQuery,
		CategoryName: asp.CategoryName,

		OrderBy: orderBy,
		Reverse: reverse,

		SkippedPages: skippedPages,
		PageSize:     pageSize,
	}
}
