package controller

import (
	"curs1_boilerplate/middleware"
	"curs1_boilerplate/util"
	"curs1_boilerplate/views/base"
	"curs1_boilerplate/views/components/navbar"
	"curs1_boilerplate/views/pages"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GeneralRestController struct {
	jwtUtil util.JwtUtil
}

func NewGeneralRestController(jwtUtil util.JwtUtil) *GeneralRestController {
	return &GeneralRestController{jwtUtil: jwtUtil}
}

func (rc *GeneralRestController) SetupRoutes(r chi.Router) {
	r.Get("/", rc.homePage)

	r.With(middleware.JwtAuth(rc.jwtUtil)).Get("/secret", rc.secretPage)
}

func (rc *GeneralRestController) homePage(w http.ResponseWriter, r *http.Request) {
	mainNavbar := navbar.MakeStandardNavbar()

	base.PageSkeleton(pages.HomePage(mainNavbar)).Render(r.Context(), w)
}

func (rc *GeneralRestController) secretPage(w http.ResponseWriter, r *http.Request) {
	base.PageSkeleton(pages.SecretPage(r.Context().Value("userEmail").(string))).Render(r.Context(), w)
}
