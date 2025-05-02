package service

type AuctionDTO struct {
	ProductName   string   `json:"productName"`
	ProductDesc   string   `json:"productDesc"`
	Category      string   `json:"category"`
	Status        string   `json:"status"`
	Mode          string   `json:"mode"`
	StartingPrice *float32 `json:"startingPrice"`
	TargetPrice   *float32 `json:"targetPrice"`
}
