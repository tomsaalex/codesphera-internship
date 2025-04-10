package infrastructure

import (
	"curs1_boilerplate/cmd/auction_rest_api/model"
	"curs1_boilerplate/db"
)

type EntityMapperDB struct {
}

func (d *EntityMapperDB) DBProductToProduct(dbProd db.Product) *model.Product {
	prod := model.Product{}

	prod.Name = dbProd.Name
	prod.Description = dbProd.Description
	prod.Price = &dbProd.Price
	prod.IsSold = dbProd.Issold

	return &prod
}

func (d *EntityMapperDB) DBProductSliceToProduct(dbProducts []db.Product) []*model.Product {
	products := make([]*model.Product, len(dbProducts))

	for i, dbProd := range dbProducts {
		products[i] = d.DBProductToProduct(dbProd)
	}

	return products
}

func (d *EntityMapperDB) ProductToCreateProductParams(product model.Product) db.CreateProductParams {
	return db.CreateProductParams{
		Name:        product.Name,
		Description: product.Description,
		Price:       *product.Price,
	}
}

func (d *EntityMapperDB) ProductToUpdateProductParams(product model.Product) db.UpdateProductParams {
	return db.UpdateProductParams{
		Name:        product.Name,
		Description: product.Description,
		Price:       *product.Price,
		Issold:      product.IsSold,
	}
}
