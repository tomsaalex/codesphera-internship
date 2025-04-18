package service

import (
	"context"
	"curs1_boilerplate/infrastructure"
	"curs1_boilerplate/model"
	"curs1_boilerplate/sharederrors"
	"curs1_boilerplate/util"
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
	ve := NewValidationError()
	validationSuccessful := true

	if userDTO.Fullname == "" {
		ve.fieldErrors["fullname"] = EMPTY
		validationSuccessful = false
	}

	if userDTO.Email == "" {
		ve.fieldErrors["email"] = EMPTY
		validationSuccessful = false
	}

	if userDTO.Password == "" {
		ve.fieldErrors["password"] = EMPTY
		validationSuccessful = false
	}

	if userDTO.Password == "" {
		ve.fieldErrors["confirmPassword"] = EMPTY
		validationSuccessful = false
	}

	if userDTO.Password != userDTO.ConfirmPassword {
		ve.fieldErrors["confirmPassword"] = INVALID
		validationSuccessful = false
	}

	if validationSuccessful {
		return nil
	}

	return ve
}

func (s *UserService) validateUserLoginDTO(userDTO UserLoginDTO) error {
	ve := NewValidationError()
	validationSuccessful := true

	if userDTO.Email == "" {
		ve.fieldErrors["email"] = EMPTY
		validationSuccessful = false
	}

	if userDTO.Password == "" {
		ve.fieldErrors["password"] = EMPTY
		validationSuccessful = false
	}

	if validationSuccessful {
		return nil
	}

	return ve
}

func (s *UserService) Register(ctx context.Context, userDTO UserRegistrationDTO) (*model.User, error) {
	err := s.validateUserRegistrationDTO(userDTO)
	if err != nil {
		return nil, err
	}

	_, err = s.userRepo.GetUserByEmail(ctx, userDTO.Email)

	if err == nil {
		return nil, &sharederrors.DuplicateEntityError{Message: "there's already a user using that email address"}
	}

	hashsalt, err := s.argonHelper.GenerateHash([]byte(userDTO.Password), nil)

	if err != nil {
		return nil, &AuthError{Message: "failed to generate hash for user's pasword"}
	}

	newUser := s.dtoMapper.RegistrationDTOToUser(userDTO, hashsalt)
	_, err = s.userRepo.Add(ctx, newUser)
	return &newUser, err
}

func (s *UserService) Login(ctx context.Context, userDTO UserLoginDTO) (*model.User, error) {
	err := s.validateUserLoginDTO(userDTO)
	if err != nil {
		return nil, err
	}

	foundUser, err := s.userRepo.GetUserByEmail(ctx, userDTO.Email)
	if err != nil {
		return nil, err
	}

	err = s.argonHelper.Compare(foundUser.PassHash, foundUser.PassSalt, []byte(userDTO.Password))

	if err != nil {
		return nil, &AuthError{Message: "auth data is incorrect"}
	}

	return foundUser, nil
}
