package infrastructure

type AuctionOrderParameter string

const (
	AuctionOrderName      = "product_name"
	AuctionOrderCreatedAt = "created_at"
)

type AuctionFilter struct {
	ProductName  string
	ProductDesc  string
	CategoryName string

	Reverse bool
	OrderBy AuctionOrderParameter

	SkippedPages int
	PageSize     int
}
