package main

type DTOMapper struct {
}

func (DTOMapper) productToSimpleProductDTO(p Product) SimpleProductDTO {
	sp := SimpleProductDTO{}

	sp.Name = p.Name
	sp.Price = p.Price

	return sp
}
