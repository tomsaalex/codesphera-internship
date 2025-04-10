package infrastructure

import (
	"context"
	"curs1_boilerplate/cmd/auction_rest_api/model"
	"fmt"
)

type MemoryProductRepo struct {
	products map[string]model.Product
}

func (r *MemoryProductRepo) GetOne(ctx context.Context, name string) (*model.Product, error) {
	prod, exists := r.products[name]
	if !exists {
		return nil, &RepositoryError{Message: fmt.Sprintf("Get One: No product matches name \"%s\"", name)}
	}

	return &prod, nil
}

func (r *MemoryProductRepo) GetAll(ctx context.Context) ([]*model.Product, error) {
	prodSlice := make([]*model.Product, len(r.products))
	i := 0
	for _, prod := range r.products {
		prodSlice[i] = &prod
		i++
	}

	return prodSlice, nil
}

func (r *MemoryProductRepo) Add(ctx context.Context, prod model.Product) (*model.Product, error) {
	_, exists := r.products[prod.Name]
	if exists {
		return nil, &RepositoryError{Message: fmt.Sprintf("Add: Name \"%s\" is already taken by a different product.", prod.Name)}
	}

	r.products[prod.Name] = prod

	return &prod, nil
}

func (r *MemoryProductRepo) Update(ctx context.Context, prod model.Product) (*model.Product, error) {
	_, exists := r.products[prod.Name]
	if !exists {
		return nil, &RepositoryError{Message: fmt.Sprintf("Update: No product matches name \"%s\"", prod.Name)}
	}

	r.products[prod.Name] = prod

	return &prod, nil
}

func (r *MemoryProductRepo) Delete(ctx context.Context, name string) error {
	_, exists := r.products[name]

	if !exists {
		return &RepositoryError{Message: fmt.Sprintf("Delete: No product matches name \"%s\"", name)}
	}

	delete(r.products, name)

	return nil
}
