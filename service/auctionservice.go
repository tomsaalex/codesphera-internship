package service

import (
	"context"
	"curs1_boilerplate/infrastructure"
	"curs1_boilerplate/middleware"
	"curs1_boilerplate/model"
	"curs1_boilerplate/sharederrors"
	"log"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type AuctionService struct {
	auctionRepo infrastructure.AuctionRepository
	userRepo    infrastructure.UserRepository
	dtoMapper   ServiceDTOMapper

	auctionCategories []model.Category
}

func NewAuctionService(auctionRepo infrastructure.AuctionRepository, userRepo infrastructure.UserRepository, dtoMapper ServiceDTOMapper) *AuctionService {
	logger := middleware.LoggerFromContext(context.Background()).With(slog.String("Layer", "NewAuctionService"))
	service := &AuctionService{
		auctionRepo: auctionRepo,
		userRepo:    userRepo,
		dtoMapper:   dtoMapper,
	}

	categories, err := service.refreshCategories()
	if err != nil {
		log.Fatal("Couldn't load categories from database")
	}

	service.auctionCategories = categories
	logger.Info("Successfully imported Auction categories from the DB")

	return service
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

	if auctionDTO.Category == "" {
		ve.fieldErrors["category"] = EMPTY
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
	logger := middleware.LoggerFromContext(ctx).With(slog.String("Layer", "AuctionService"))

	err := s.validateAuctionDTO(auctionDTO)
	if err != nil {
		logger.Error("Given Auction is invalid")
		return nil, err
	}

	auctionDTO = s.sanitizeAuctionDTO(auctionDTO)

	_, err = s.auctionRepo.GetAuctionByName(ctx, auctionDTO.ProductName)

	if err == nil {
		logger.Error("Auction's product name is duplicated")
		return nil, &sharederrors.DuplicateEntityError{Message: "there's already a product by that name being auctioned"}
	}
	sellerEmail := middleware.GetUserEmailFromContext(ctx)
	auctionSeller, err := s.userRepo.GetUserByEmail(ctx, sellerEmail)

	if err != nil {
		logger.Error("Auction's seller couldn't be found in the DB")
		return nil, err
	}

	newAuction, err := s.dtoMapper.AuctionDTOToAuction(auctionDTO, auctionSeller, s.auctionCategories)

	if err != nil {
		logger.Warn("Couldn't convert AuctionDTO to Auction.")
		// TODO: Make this nicer. Wrap the underlying error maybe (which in itself should be nicer).
		return nil, &ServiceError{Message: "auction fields are invalid"}
	}

	newAuction.CreatedAt = time.Now()
	savedAuction, err := s.auctionRepo.Add(ctx, *newAuction)
	return savedAuction, err
}

func (s *AuctionService) GetAuctions(ctx context.Context, auctionFilter infrastructure.AuctionFilter) ([]model.Auction, int, error) {
	return s.auctionRepo.GetAuctions(ctx, auctionFilter)
}

func (s *AuctionService) GetAuctionById(ctx context.Context, auctionId uuid.UUID) (*model.Auction, error) {
	return s.auctionRepo.GetAuctionById(ctx, auctionId)
}

func (s *AuctionService) refreshCategories() ([]model.Category, error) {
	return s.auctionRepo.GetCategories(context.Background())
}

func (s *AuctionService) GetCachedCategories() []model.Category {
	return s.auctionCategories
}
