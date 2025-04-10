package infrastructure

import (
	"context"
	"curs1_boilerplate/cmd/auction_rest_api/model"
	"curs1_boilerplate/db"
)

type ProductRepository interface {
	GetOne(context.Context, string) (*model.Product, error)
	GetAll(context.Context) ([]*model.Product, error)
	Add(context.Context, model.Product) (*model.Product, error)
	Update(context.Context, model.Product) (*model.Product, error)
	Delete(context.Context, string) error
}

func NewMemoryProductRepository() ProductRepository {
	return &MemoryProductRepo{
		products: make(map[string]model.Product),
	}
}

func NewDBProductRepository(queries *db.Queries) ProductRepository {
	return &DBProductRepo{
		queries: queries,
		mapper:  EntityMapperDB{},
	}
}
