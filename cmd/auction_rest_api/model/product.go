package model

import "encoding/json"

type Product struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       *float64 `json:"price"`
	IsSold      bool     `json:"isSold"`
}

func (p *Product) UnmarshalJSON(data []byte) error {
	type alias Product
	var aux = &struct {
		*alias
	}{
		alias: (*alias)(p),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	aux.IsSold = false

	return nil
}

type SimpleProductDTO struct {
	Name  string   `json:"name"`
	Price *float64 `json:"price"`
}
