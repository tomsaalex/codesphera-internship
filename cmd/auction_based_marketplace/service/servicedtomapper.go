package service

import (
	"curs1_boilerplate/cmd/auction_based_marketplace/model"
	"curs1_boilerplate/cmd/auction_based_marketplace/util"
)

type ServiceDTOMapper struct {
}

func (m *ServiceDTOMapper) RegistrationDTOToUser(userDTO UserRegistrationDTO, hashsalt *util.HashSalt) model.User {
	return model.User{
		Fullname: userDTO.Fullname,
		Email:    userDTO.Email,
		PassHash: hashsalt.Hash,
		PassSalt: hashsalt.Salt,
	}
}

func NewServiceDTOMapper() *ServiceDTOMapper {
	return &ServiceDTOMapper{}
}
