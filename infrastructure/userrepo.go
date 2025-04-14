package infrastructure

import (
	"context"
	"curs1_boilerplate/db"
	"curs1_boilerplate/model"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	Add(ctx context.Context, newUser model.User) (*model.User, error)
	Update(ctx context.Context, updatedUser model.User) (*model.User, error)
	Delete(ctx context.Context, userEmail string) error
}

func NewDBUserRepository(queries *db.Queries) UserRepository {
	return &DBUserRepo{
		queries: queries,
		mapper:  EntityMapperDB{},
	}
}
