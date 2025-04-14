package controller

import (
	"curs1_boilerplate/views/base"
	"curs1_boilerplate/views/components"
	"curs1_boilerplate/views/pages"
	"net/http"

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
	base.PageSkeleton(pages.HomePage(components.Navbar())).Render(r.Context(), w)
}
