package service

import (
	"curs1_boilerplate/model"
	"curs1_boilerplate/util"
	"fmt"
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

func (m *ServiceDTOMapper) AuctionDTOToAuction(auctionDTO AuctionDTO, seller *model.User, categories []model.Category) (*model.Auction, error) {
	auctionMode, err := m.stringToAuctionMode(auctionDTO.Mode)

	if err != nil {
		return nil, err
	}

	auctionStatus, err := m.stringToAuctionStatus(auctionDTO.Status)

	if err != nil {
		return nil, err
	}

	var targetPrice *float32 = nil
	if auctionDTO.TargetPrice != nil {
		targetPrice = auctionDTO.TargetPrice
	}

	var category *model.Category = nil

	for _, cat := range categories {
		if auctionDTO.Category == cat.Name {
			category = &cat
		}
	}

	if category == nil {
		return nil, fmt.Errorf("couldn't identify given category")
	}

	return &model.Auction{
		ProductName:        auctionDTO.ProductName,
		ProductDescription: auctionDTO.ProductDesc,
		Category:           category,
		Status:             auctionStatus,
		Mode:               auctionMode,
		StartingPrice:      auctionDTO.StartingPrice,
		TargetPrice:        targetPrice,
		Seller:             seller,
	}, nil
}

func (s *ServiceDTOMapper) stringToAuctionMode(rawMode string) (model.AuctionMode, error) {
	switch rawMode {
	case "Manual":
		return model.AM_Manual, nil
	case "Price Met":
		return model.AM_Price_Met, nil
	default:
		return model.AM_Manual, fmt.Errorf("couldn't convert string AuctionMode enum to model")
	}
}

func (s *ServiceDTOMapper) stringToAuctionStatus(rawStatus string) (model.AuctionStatus, error) {
	switch rawStatus {
	case "Scheduled":
		return model.AS_Scheduled, nil
	case "Immediate":
		return model.AS_Ongoing, nil
	case "Finished":
		return model.AS_Finished, nil
	default:
		return model.AS_Ongoing, fmt.Errorf("couldn't convert db AuctionStatus enum to model variant")
	}
}

func NewServiceDTOMapper() *ServiceDTOMapper {
	return &ServiceDTOMapper{}
}
