package model

type DTOMapper struct {
}

func (DTOMapper) ProductToSimpleProductDTO(p Product) SimpleProductDTO {
	sp := SimpleProductDTO{}

	sp.Name = p.Name
	sp.Price = p.Price

	return sp
}
