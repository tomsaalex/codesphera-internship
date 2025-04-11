package service

import (
	"context"
	"curs1_boilerplate/cmd/auction_based_marketplace/infrastructure"
	"curs1_boilerplate/cmd/auction_based_marketplace/util"
)

type UserService struct {
	userRepo    infrastructure.UserRepository
	dtoMapper   ServiceDTOMapper
	argonHelper util.Argon2idHash
}

func NewUserService(userRepo infrastructure.UserRepository, dtoMapper ServiceDTOMapper, argonHelper util.Argon2idHash) *UserService {
	return &UserService{
		userRepo:    userRepo,
		dtoMapper:   dtoMapper,
		argonHelper: argonHelper,
	}
}

func (s *UserService) Register(ctx context.Context, userDTO UserRegistrationDTO) error {
	_, err := s.userRepo.GetUserByEmail(ctx, userDTO.Email)

	if err == nil {
		return &ServiceError{Message: "There's already a user using that email address."}
	}

	hashsalt, err := s.argonHelper.GenerateHash([]byte(userDTO.Password), nil)

	if err != nil {
		return &ServiceError{Message: "Failed to generate hash for user's pasword."}
	}

	newUser := s.dtoMapper.RegistrationDTOToUser(userDTO, hashsalt)
	_, err = s.userRepo.Add(ctx, newUser)
	return err
}

func (s *UserService) Login(ctx context.Context, userDTO UserLoginDTO) error {
	foundUser, err := s.userRepo.GetUserByEmail(ctx, userDTO.Email)
	if err != nil {
		return &ServiceError{Message: "No user matches the given email."}
	}

	err = s.argonHelper.Compare(foundUser.PassHash, foundUser.PassSalt, []byte(userDTO.Password))

	if err != nil {
		// TODO: Not strictly true. This could also be just an error while generating the hash... but I guess that means the data is just broken.
		return &ServiceError{Message: "Auth data is incorrect"}
	}

	return nil
}
