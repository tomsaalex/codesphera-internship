package service

import (
	"curs1_boilerplate/model"
	"curs1_boilerplate/util"
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
