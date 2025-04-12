package service

import (
	"context"
	"curs1_boilerplate/cmd/auction_based_marketplace/infrastructure"
	"curs1_boilerplate/cmd/auction_based_marketplace/sharederrors"
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

func (s *UserService) validateUserRegistrationDTO(userDTO UserRegistrationDTO) error {
	reqerrs := ""

	if userDTO.Fullname == "" {
		reqerrs += "Cannot register a user with no name.\n"
	}

	if userDTO.Email == "" {
		reqerrs += "Cannot register a user with no email.\n"
	}

	if userDTO.Password == "" {
		reqerrs += "Cannot register a user with no password.\n"
	}

	if reqerrs != "" {
		return &ValidationError{Message: reqerrs}
	}

	return nil
}

func (s *UserService) validateUserLoginDTO(userDTO UserLoginDTO) error {
	reqerrs := ""

	if userDTO.Email == "" {
		reqerrs += "Users can't have a blank email.\n"
	}

	if userDTO.Password == "" {
		reqerrs += "Users can't have a blank password.\n"
	}

	if reqerrs != "" {
		return &ValidationError{Message: reqerrs}
	}

	return nil
}

func (s *UserService) Register(ctx context.Context, userDTO UserRegistrationDTO) error {
	err := s.validateUserRegistrationDTO(userDTO)
	if err != nil {
		return err
	}

	_, err = s.userRepo.GetUserByEmail(ctx, userDTO.Email)

	if err == nil {
		return &sharederrors.DuplicateEntityError{Message: "there's already a user using that email address"}
	}

	hashsalt, err := s.argonHelper.GenerateHash([]byte(userDTO.Password), nil)

	if err != nil {
		return &AuthError{Message: "failed to generate hash for user's pasword"}
	}

	newUser := s.dtoMapper.RegistrationDTOToUser(userDTO, hashsalt)
	_, err = s.userRepo.Add(ctx, newUser)
	return err
}

func (s *UserService) Login(ctx context.Context, userDTO UserLoginDTO) error {
	err := s.validateUserLoginDTO(userDTO)
	if err != nil {
		return err
	}

	foundUser, err := s.userRepo.GetUserByEmail(ctx, userDTO.Email)
	if err != nil {
		return err
	}

	err = s.argonHelper.Compare(foundUser.PassHash, foundUser.PassSalt, []byte(userDTO.Password))

	if err != nil {
		return &AuthError{Message: "auth data is incorrect"}
	}

	return nil
}
