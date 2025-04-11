package infrastructure

import (
	"curs1_boilerplate/cmd/auction_based_marketplace/model"
	"curs1_boilerplate/db"
)

type EntityMapperDB struct {
}

func (d *EntityMapperDB) DBUserToUser(dbUser db.User) *model.User {
	user := model.User{
		Id:       dbUser.ID.Bytes,
		Email:    dbUser.Email,
		Fullname: dbUser.Fullname,
		PassHash: dbUser.PassHash,
		PassSalt: dbUser.PassSalt,
	}

	return &user
}

func (d *EntityMapperDB) UserToAddUserParams(user model.User) db.AddUserParams {
	return db.AddUserParams{
		Fullname: user.Fullname,
		Email:    user.Email,
		PassHash: user.PassHash,
		PassSalt: user.PassSalt,
	}
}

func (d *EntityMapperDB) UserToUpdateUserParams(user model.User) db.UpdateUserParams {
	return db.UpdateUserParams{
		Fullname: user.Fullname,
		Email:    user.Email,
		PassHash: user.PassHash,
		PassSalt: user.PassSalt,
	}
}
