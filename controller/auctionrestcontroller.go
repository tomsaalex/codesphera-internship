package controller

import (
	"curs1_boilerplate/infrastructure"
	"curs1_boilerplate/middleware"
	"curs1_boilerplate/service"
	"curs1_boilerplate/sharederrors"
	"curs1_boilerplate/util"
	"curs1_boilerplate/views/base"
	"curs1_boilerplate/views/components/navbar"
	"curs1_boilerplate/views/pages/aucbrowse"
	"curs1_boilerplate/views/pages/auccreate"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuctionRestController struct {
	auctionService service.AuctionService
	jwtHelper      util.JwtUtil
}

func NewAuctionRestController(auctionService service.AuctionService, jwtHelper util.JwtUtil) *AuctionRestController {
	return &AuctionRestController{
		auctionService: auctionService,
		jwtHelper:      jwtHelper,
	}
}

func (rc *AuctionRestController) SetupRoutes(r chi.Router) {
	r.With(middleware.RequireAuth).Get("/start-auction", rc.createAuctionPage)
	r.With(middleware.RequireAuth).Post("/auctions", rc.addAuction)
	r.Get("/search-auctions", rc.browseAuctions)
}

func (rc *AuctionRestController) createAuctionPage(w http.ResponseWriter, r *http.Request) {
	createAuctionPage := auccreate.MakeValidAuctionCreationPage(false, navbar.MakeStandardNavbar(r.Context()))

	base.PageSkeleton(createAuctionPage).Render(r.Context(), w)
}

func (rc *AuctionRestController) browseAuctions(w http.ResponseWriter, r *http.Request) {
	auctions, _ := rc.auctionService.GetAuctions(r.Context())
	browseAuctionsPage := aucbrowse.MakeAuctionBrowsePage(auctions, navbar.MakeStandardNavbar(r.Context()))
	base.PageSkeleton(browseAuctionsPage).Render(r.Context(), w)
}

func (rc *AuctionRestController) addAuction(w http.ResponseWriter, r *http.Request) {
	var auctionDTO service.AuctionDTO
	err := json.NewDecoder(r.Body).Decode(&auctionDTO)
	formErrs := auccreate.AuctionCreateFormErrors{}
	if err != nil {
		formErrs.GenericError = "Request failed for an unknown reason"

		auctionCreationPage := auccreate.MakeErroredAuctionCreationPage(&formErrs, false, navbar.MakeStandardNavbar(r.Context()))
		base.PageSkeleton(auctionCreationPage).Render(r.Context(), w)
		return
	}

	// TODO: I despise using raw strings like this with my entire being, but it'll have to do for now
	formHasTargetPrice := auctionDTO.Mode == "Price Met"

	_, err = rc.auctionService.AddAuction(r.Context(), auctionDTO)
	if err != nil {
		var entityNotFoundErr *infrastructure.EntityNotFoundError
		var valErr *service.ValidationError
		var duplicateErr *sharederrors.DuplicateEntityError
		var serviceErr *service.ServiceError
		var repositoryErr *infrastructure.RepositoryError
		var entityDBMappingErr *infrastructure.EntityDBMappingError
		var foreignKeyViolationErr *infrastructure.ForeignKeyViolationError

		var auctionCreationPage *auccreate.ViewModel

		if errors.As(err, &valErr) {
			_, hasProductNameErr := valErr.GetField("productName")
			_, hasProductDescErr := valErr.GetField("productDesc")
			_, hasModeErr := valErr.GetField("mode")
			_, hasStatusErr := valErr.GetField("status")
			startingPriceErr, hasStartingPriceErr := valErr.GetField("startingPrice")
			targetPriceErr, hasTargetPriceErr := valErr.GetField("targetPrice")

			if hasProductNameErr {
				formErrs.ProductNameError = "You can't create an auction with a blank product name."
			}
			if hasProductDescErr {
				formErrs.ProductDescError = "You can't create an auction with a blank product description."
			}
			if hasModeErr {
				formErrs.ModeError = "You can't create an auction without specifying the auction mode."
			}
			if hasStatusErr {
				formErrs.StatusError = "You can't create an auction without specifying its status."
			}
			if hasStartingPriceErr {
				switch startingPriceErr {
				case service.EMPTY:
					formErrs.StartingPriceError = "You can't create an auction without specifying a starting price."
				case service.NEGATIVE:
					formErrs.StartingPriceError = "You can't create an auction with a negative starting price."
				}
			}
			if hasTargetPriceErr {
				switch targetPriceErr {
				case service.EMPTY:
					formErrs.TargetPriceError = "You can't create an auction without specifying a target price if the auction mode is set to Price Met."
				case service.NEGATIVE:
					formErrs.TargetPriceError = "You can't create an auction with a negative target price."
				case service.INVALID:
					formErrs.TargetPriceError = "You can't create an auction that has the target price lower than the starting price."
				}
			}
		} else if errors.As(err, &duplicateErr) {
			formErrs.ProductNameError = "There's already an auction for a product with the same name."
		} else if errors.As(err, &entityNotFoundErr) {
			formErrs.GenericError = "Couldn't create an auction for the given seller. Seller hasn't been found."
		} else if errors.As(err, &serviceErr) {
			formErrs.GenericError = "Auction fields are invalid."
		} else if errors.As(err, &entityDBMappingErr) {
			formErrs.GenericError = "Auction fields are invalid."
		} else if errors.As(err, &foreignKeyViolationErr) {
			formErrs.GenericError = "Couldn't create an auction for the given seller. Seller hasn't been found."
		} else if errors.As(err, &repositoryErr) {
			formErrs.GenericError = "Auction fields are invalid."
		} else {
			formErrs.GenericError = "An unexpected error occurred on our end. Please retry later!"
		}
		auctionCreationPage = auccreate.MakeErroredAuctionCreationPage(&formErrs, formHasTargetPrice, navbar.MakeStandardNavbar(r.Context()))
		base.PageSkeleton(auctionCreationPage).Render(r.Context(), w)
		return
	}

	// TODO: Redirect to auction page... after you make it
	w.Header().Set("HX-Redirect", "/")
}
