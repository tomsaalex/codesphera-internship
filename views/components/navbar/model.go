package navbar

import (
	"context"
	"curs1_boilerplate/views/components/anchor"
	"curs1_boilerplate/views/components/buttongroup"
	"curs1_boilerplate/views/components/searchbar"
	"io"

	"github.com/a-h/templ"
)

type Model struct {
	id         string
	links      map[string]string
	searchbar  templ.Component
	authButton templ.Component
}

func Make(id string, links map[string]string, searchbar, authButton templ.Component) *Model {
	return &Model{
		id:         id,
		links:      links,
		searchbar:  searchbar,
		authButton: authButton,
	}
}

func MakeStandardNavbar() templ.Component {
	navLinks := make(map[string]string)
	navLinks["Home"] = "/"
	navLinks["About Us"] = "#"
	navLinks["Start An Auction"] = "#"

	navSearch := searchbar.Make("nav-search", "Search for auctions", "Search auctions")

	registerButton := anchor.Make("register-button", "Register!", "/register")
	loginButton := anchor.Make("login-button", "Log in!", "/login")

	navAuthButton := buttongroup.Make("nav-auth", "Login/Register Buttons", []templ.Component{loginButton, registerButton})
	return Make("main-nav", navLinks, navSearch, navAuthButton)
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return component(m).Render(ctx, w)
}
