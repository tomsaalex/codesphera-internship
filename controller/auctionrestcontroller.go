package controller

import (
	"curs1_boilerplate/middleware"
	"curs1_boilerplate/service"
	"curs1_boilerplate/util"
	"curs1_boilerplate/views/base"
	"curs1_boilerplate/views/components/navbar"
	"curs1_boilerplate/views/pages/auccreate"
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
}

func (rc *AuctionRestController) createAuctionPage(w http.ResponseWriter, r *http.Request) {
	createAuctionPage := auccreate.MakeValidAuctionCreationPage(navbar.MakeStandardNavbar(r.Context()))

	base.PageSkeleton(createAuctionPage).Render(r.Context(), w)
}
