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

func (m *ServiceDTOMapper) AuctionDTOToAuction(auctionDTO AuctionDTO, seller *model.User) (*model.Auction, error) {
	auctionMode, err := m.stringToAuctionMode(auctionDTO.Mode)

	if err != nil {
		return nil, err
	}

	auctionStatus, err := m.stringToAuctionStatus(auctionDTO.Status)

	if err != nil {
		return nil, err
	}

	return &model.Auction{
		ProductName:        auctionDTO.ProductName,
		ProductDescription: auctionDTO.ProductDesc,
		Status:             auctionStatus,
		Mode:               auctionMode,
		StartingPrice:      *auctionDTO.StartingPrice,
		TargetPrice:        *auctionDTO.TargetPrice,
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
		// TODO: Reconsider whether this is necessary... this case should be caught by validator.. but still
		// Make the error nicer, too
		return model.AM_Manual, fmt.Errorf("couldn't convert string AuctionMode enum to model")
	}
}

func (s *ServiceDTOMapper) stringToAuctionStatus(rawStatus string) (model.AuctionStatus, error) {
	switch rawStatus {
	case "Ongoing":
		return model.AS_Ongoing, nil
	case "Finished":
		return model.AS_Finished, nil
	default:
		// TODO: Reconsider whether this is necessary... this case should be caught by validator.. but still
		// Make the error nicer, too
		return model.AS_Ongoing, fmt.Errorf("couldn't convert db AuctionStatus enum to model variant")
	}
}

func NewServiceDTOMapper() *ServiceDTOMapper {
	return &ServiceDTOMapper{}
}
