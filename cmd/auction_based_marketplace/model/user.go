package model

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Fullname string
	Email    string
	PassHash []byte
	PassSalt []byte
}
