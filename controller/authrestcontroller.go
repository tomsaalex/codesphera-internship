package controller

import (
	"curs1_boilerplate/middleware"
	"curs1_boilerplate/views/base"
	"curs1_boilerplate/views/components/navbar"
	loginpage "curs1_boilerplate/views/pages/login"
	registerpage "curs1_boilerplate/views/pages/register"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthRestController struct {
}

func NewAuthRestController() *AuthRestController {
	return &AuthRestController{}
}

func (rc *AuthRestController) SetupRoutes(r chi.Router) {
	r.With(middleware.RequireGuest).Get("/login", rc.loginPage)
	r.With(middleware.RequireGuest).Get("/register", rc.registerPage)
}

func (rc *AuthRestController) loginPage(w http.ResponseWriter, r *http.Request) {
	loginPage := loginpage.MakeValidLoginPage(navbar.MakeStandardNavbar(r.Context()))

	base.PageSkeleton(loginPage).Render(r.Context(), w)
}

func (rc *AuthRestController) registerPage(w http.ResponseWriter, r *http.Request) {
	base.PageSkeleton(registerpage.MakeValidRegisterPage(navbar.MakeStandardNavbar(r.Context()))).Render(r.Context(), w)
}
