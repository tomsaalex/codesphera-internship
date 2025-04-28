package service

import (
	"context"
	"curs1_boilerplate/infrastructure"
	"curs1_boilerplate/middleware"
	"curs1_boilerplate/model"
	"curs1_boilerplate/sharederrors"
)

type AuctionService struct {
	auctionRepo infrastructure.AuctionRepository
	userRepo    infrastructure.UserRepository
	dtoMapper   ServiceDTOMapper
}

func NewAuctionService(auctionRepo infrastructure.AuctionRepository, userRepo infrastructure.UserRepository, dtoMapper ServiceDTOMapper) *AuctionService {
	return &AuctionService{
		auctionRepo: auctionRepo,
		userRepo:    userRepo,
		dtoMapper:   dtoMapper,
	}
}

func (s *AuctionService) sanitizeAuctionDTO(auctionDTO AuctionDTO) AuctionDTO {
	if auctionDTO.Mode != "Price Met" {
		auctionDTO.TargetPrice = nil
	}

	return auctionDTO
}

func (s *AuctionService) validateAuctionDTO(auctionDTO AuctionDTO) error {
	ve := NewValidationError()
	validationSuccessful := true

	if auctionDTO.ProductName == "" {
		ve.fieldErrors["productName"] = EMPTY
		validationSuccessful = false
	}

	if auctionDTO.ProductDesc == "" {
		ve.fieldErrors["productDesc"] = EMPTY
		validationSuccessful = false
	}

	if auctionDTO.Mode == "" {
		ve.fieldErrors["mode"] = EMPTY
		validationSuccessful = false
	}

	if auctionDTO.Status == "" {
		ve.fieldErrors["status"] = EMPTY
		validationSuccessful = false
	}

	if auctionDTO.StartingPrice == nil {
		ve.fieldErrors["startingPrice"] = EMPTY
		validationSuccessful = false
	}

	if auctionDTO.StartingPrice != nil && *auctionDTO.StartingPrice < 0 {
		ve.fieldErrors["startingPrice"] = NEGATIVE
		validationSuccessful = false
	}

	if auctionDTO.Mode == "Price Met" {
		if auctionDTO.TargetPrice == nil {
			ve.fieldErrors["targetPrice"] = EMPTY
			validationSuccessful = false
		}

		if auctionDTO.TargetPrice != nil && *auctionDTO.TargetPrice < 0 {
			ve.fieldErrors["targetPrice"] = NEGATIVE
			validationSuccessful = false
		}

		if auctionDTO.TargetPrice != nil && *auctionDTO.TargetPrice < *auctionDTO.StartingPrice {
			ve.fieldErrors["targetPrice"] = INVALID
			validationSuccessful = false
		}
	}
	if validationSuccessful {
		return nil
	}

	return ve
}

func (s *AuctionService) AddAuction(ctx context.Context, auctionDTO AuctionDTO) (*model.Auction, error) {
	err := s.validateAuctionDTO(auctionDTO)
	if err != nil {
		return nil, err
	}

	auctionDTO = s.sanitizeAuctionDTO(auctionDTO)

	_, err = s.auctionRepo.GetAuctionByName(ctx, auctionDTO.ProductName)

	if err == nil {
		return nil, &sharederrors.DuplicateEntityError{Message: "there's already a product by that name being auctioned"}
	}
	sellerEmail := middleware.GetUserEmailFromContext(ctx)
	auctionSeller, err := s.userRepo.GetUserByEmail(ctx, sellerEmail)

	if err != nil {
		return nil, err
	}

	newAuction, err := s.dtoMapper.AuctionDTOToAuction(auctionDTO, auctionSeller)

	if err != nil {
		// TODO: Make this nicer. Wrap the underlying error maybe (which in itself should be nicer).
		return nil, &ServiceError{Message: "auction fields are invalid"}
	}
	savedAuction, err := s.auctionRepo.Add(ctx, *newAuction)
	return savedAuction, err
}

func (s *AuctionService) GetAuctions(ctx context.Context) ([]model.Auction, error) {
	return s.auctionRepo.GetAuctions(ctx)
}
