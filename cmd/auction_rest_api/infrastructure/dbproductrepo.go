package infrastructure

import (
	"context"
	"curs1_boilerplate/cmd/auction_rest_api/model"
	"curs1_boilerplate/db"
	"fmt"
)

type DBProductRepo struct {
	queries *db.Queries
	mapper  EntityMapperDB
}

func (r *DBProductRepo) GetOne(ctx context.Context, name string) (*model.Product, error) {
	foundProduct, err := r.queries.GetProduct(ctx, name)
	if err != nil {
		return nil, &RepositoryError{Message: fmt.Sprintf("Get One: No product matches name \"%s\"", name)}
	}

	convertedProduct := r.mapper.DBProductToProduct(foundProduct)
	return convertedProduct, nil
}

func (r *DBProductRepo) GetAll(ctx context.Context) ([]*model.Product, error) {
	products, err := r.queries.ListProducts(ctx)
	if err != nil {
		return nil, &RepositoryError{Message: "Unknown error occured while fetching products from database."}
	}

	convertedProducts := r.mapper.DBProductSliceToProduct(products)
	return convertedProducts, nil
}

func (r *DBProductRepo) Add(ctx context.Context, prod model.Product) (*model.Product, error) {
	createProductParams := r.mapper.ProductToCreateProductParams(prod)
	dbProduct, err := r.queries.CreateProduct(ctx, createProductParams)
	if err != nil {
		return nil, &RepositoryError{Message: fmt.Sprintf("Add: Name \"%s\" is already taken by a different product.", prod.Name)}
	}

	modelProduct := r.mapper.DBProductToProduct(dbProduct)
	return modelProduct, nil
}

func (r *DBProductRepo) Update(ctx context.Context, prod model.Product) (*model.Product, error) {
	updateProductParams := r.mapper.ProductToUpdateProductParams(prod)
	dbProduct, err := r.queries.UpdateProduct(ctx, updateProductParams)

	if err != nil {
		return nil, &RepositoryError{Message: fmt.Sprintf("Update: No product matches name \"%s\"", prod.Name)}
	}

	modelProduct := r.mapper.DBProductToProduct(dbProduct)
	return modelProduct, nil
}

func (r *DBProductRepo) Delete(ctx context.Context, name string) error {
	err := r.queries.DeleteProduct(ctx, name)

	if err != nil {
		return &RepositoryError{Message: fmt.Sprintf("Delete: No product matches name \"%s\"", name)}
	}

	return nil
}
