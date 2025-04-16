package controller

import (
	"curs1_boilerplate/views/base"
	"curs1_boilerplate/views/components/navbar"
	"curs1_boilerplate/views/pages"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthRestController struct {
}

func NewAuthRestController() *AuthRestController {
	return &AuthRestController{}
}

func (rc *AuthRestController) SetupRoutes(r chi.Router) {
	r.Get("/login", rc.loginPage)
	r.Get("/register", rc.registerPage)
}

func (rc *AuthRestController) loginPage(w http.ResponseWriter, r *http.Request) {
	base.PageSkeleton(pages.LoginPage(navbar.MakeStandardNavbar())).Render(r.Context(), w)
}

func (rc *AuthRestController) registerPage(w http.ResponseWriter, r *http.Request) {
	base.PageSkeleton(pages.RegisterPage(navbar.MakeStandardNavbar())).Render(r.Context(), w)
}
