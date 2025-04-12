package infrastructure

import (
	"context"
	"curs1_boilerplate/cmd/auction_based_marketplace/model"
	"curs1_boilerplate/cmd/auction_based_marketplace/sharederrors"
	"curs1_boilerplate/db"
	"fmt"
)

type DBUserRepo struct {
	queries *db.Queries
	mapper  EntityMapperDB
}

func (r *DBUserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	foundUser, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, &EntityNotFoundError{Message: fmt.Sprintf("GetUserByEmail: No user matches email \"%s\"", email)}
	}

	convertedUser := r.mapper.DBUserToUser(foundUser)
	return convertedUser, nil
}

func (r *DBUserRepo) Add(ctx context.Context, user model.User) (*model.User, error) {
	addUserParams := r.mapper.UserToAddUserParams(user)
	dbUser, err := r.queries.AddUser(ctx, addUserParams)
	if err != nil {
		return nil, &sharederrors.DuplicateEntityError{Message: fmt.Sprintf("Add: Email \"%s\" is already taken by a different user.", user.Email)}
	}

	modelUser := r.mapper.DBUserToUser(dbUser)
	return modelUser, nil
}

func (r *DBUserRepo) Update(ctx context.Context, user model.User) (*model.User, error) {
	updateUserParams := r.mapper.UserToUpdateUserParams(user)
	dbUser, err := r.queries.UpdateUser(ctx, updateUserParams)

	if err != nil {
		return nil, &EntityNotFoundError{Message: fmt.Sprintf("Update: No user matches email \"%s\"", user.Email)}
	}

	modelUser := r.mapper.DBUserToUser(dbUser)
	return modelUser, nil
}

func (r *DBUserRepo) Delete(ctx context.Context, email string) error {
	dbUser, err := r.queries.GetUserByEmail(ctx, email)

	if err != nil {
		return &EntityNotFoundError{Message: fmt.Sprintf("Delete: No user matches email \"%s\"", email)}
	}

	err = r.queries.DeleteUser(ctx, dbUser.ID)

	if err != nil {
		return &RepositoryError{Message: fmt.Sprintf("Delete: failed to delete user with email \"%s\"", email)}
	}

	return nil
}
