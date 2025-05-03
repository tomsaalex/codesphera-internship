package navbar

import (
	"context"
	"curs1_boilerplate/middleware"
	"curs1_boilerplate/views/components/anchor"
	"curs1_boilerplate/views/components/buttongroup"
	profilebutton "curs1_boilerplate/views/components/profile_button"
	"curs1_boilerplate/views/components/searchbar"
	"io"

	"github.com/a-h/templ"
)

type NavLink struct {
	LinkText        string
	LinkDestination string
}

type Model struct {
	id         string
	links      []NavLink
	searchbar  templ.Component
	authButton templ.Component
}

func Make(id string, links []NavLink, searchbar, authButton templ.Component) *Model {
	return &Model{
		id:         id,
		links:      links,
		searchbar:  searchbar,
		authButton: authButton,
	}
}

func MakeStandardNavLinks() []NavLink {
	navLinks := make([]NavLink, 0)
	navLinks = append(navLinks, NavLink{LinkText: "Home", LinkDestination: "/"})
	navLinks = append(navLinks, NavLink{LinkText: "About Us", LinkDestination: "#"})
	navLinks = append(navLinks, NavLink{LinkText: "Start An Auction", LinkDestination: "/start-auction"})
	navLinks = append(navLinks, NavLink{LinkText: "Browse Auctions", LinkDestination: "/search-auctions"})
	return navLinks
}

func MakeStandardNavbarCustomSearch(ctx context.Context, searchbar searchbar.Model) *Model {
	navLinks := MakeStandardNavLinks()

	userEmail := middleware.GetUserEmailFromContext(ctx)

	var navbarAuthComp templ.Component

	if userEmail != "" {
		navbarAuthComp = profilebutton.Make(userEmail)
	} else {
		registerButton := anchor.Make("register-button", "Register!", "/register")
		loginButton := anchor.Make("login-button", "Log in!", "/login")

		navbarAuthComp = buttongroup.Make("nav-auth", "Login/Register Buttons", []templ.Component{loginButton, registerButton})
	}

	return Make("main-nav", navLinks, &searchbar, navbarAuthComp)
}

// TODO: Replace context with profile info, maybe
func MakeStandardNavbar(ctx context.Context) *Model {
	navSearch := searchbar.Make("nav-search", "Search for auctions", "Search auctions", "home")
	return MakeStandardNavbarCustomSearch(ctx, *navSearch)
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return component(m).Render(ctx, w)
}
