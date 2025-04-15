package controller

import (
	"curs1_boilerplate/views/base"
	"curs1_boilerplate/views/components/anchor"
	"curs1_boilerplate/views/components/buttongroup"
	"curs1_boilerplate/views/components/navbar"
	"curs1_boilerplate/views/components/searchbar"
	"curs1_boilerplate/views/pages"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

type GeneralRestController struct {
}

func NewGeneralRestController() *GeneralRestController {
	return &GeneralRestController{}
}

func (rc *GeneralRestController) SetupRoutes(r chi.Router) {
	r.Get("/", rc.homePage)
}

func (rc *GeneralRestController) homePage(w http.ResponseWriter, r *http.Request) {
	// TODO: Change this, Go maps don't preserve order
	navLinks := make(map[string]string)
	navLinks["Home"] = "/"
	navLinks["About Us"] = "#"
	navLinks["Start An Auction"] = "#"

	navSearch := searchbar.Make("nav-search", "Search for auctions", "Search auctions")

	registerButton := anchor.Make("register-button", "Register!", "/register")
	loginButton := anchor.Make("login-button", "Log in!", "/login")

	navAuthButton := buttongroup.Make("nav-auth", "Login/Register Buttons", []templ.Component{loginButton, registerButton})
	mainNavbar := navbar.Make("main-nav", navLinks, navSearch, navAuthButton)
	base.PageSkeleton(pages.HomePage(mainNavbar)).Render(r.Context(), w)
}
