package controller

import (
	"curs1_boilerplate/infrastructure"
	"curs1_boilerplate/service"
	"curs1_boilerplate/sharederrors"
	"curs1_boilerplate/util"
	"curs1_boilerplate/views/base"
	"curs1_boilerplate/views/components/navbar"
	loginpage "curs1_boilerplate/views/pages/login"
	registerpage "curs1_boilerplate/views/pages/register"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserRestController struct {
	userService service.UserService
	jwtHelper   util.JwtUtil
}

func NewUserRestController(userService service.UserService, jwtHelper util.JwtUtil) *UserRestController {
	return &UserRestController{
		userService: userService,
		jwtHelper:   jwtHelper,
	}
}

func (rc *UserRestController) SetupRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", rc.registerUser)
		r.Post("/login", rc.loginUser)
		r.Get("/logout", rc.logoutUser)
	})
	r.Route("/user", func(r chi.Router) {
		r.Delete("/", rc.deleteUser)
		r.Put("/", rc.editUser)
	})
}

func (rc *UserRestController) logoutUser(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "authCookie",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	w.Header().Set("HX-Redirect", "/")
}

func (rc *UserRestController) registerUser(w http.ResponseWriter, r *http.Request) {
	var userDTO service.UserRegistrationDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)

	formErrs := registerpage.RegisterFormErrors{}
	if err != nil {
		formErrs.GenericError = "Request failed for an unknown reason"

		registerPage := registerpage.MakeErroredRegisterPage(&formErrs, navbar.MakeStandardNavbar(r.Context()))
		base.PageSkeleton(registerPage).Render(r.Context(), w)
		return
	}

	registeredUser, err := rc.userService.Register(r.Context(), userDTO)
	if err != nil {
		var authErr *service.AuthError
		var duplicateEntityErr *sharederrors.DuplicateEntityError
		var valErr *service.ValidationError

		if errors.As(err, &authErr) {
			formErrs.PasswordError = "Password could not be used. Please try again with another password."
		} else if errors.As(err, &duplicateEntityErr) {
			formErrs.EmailError = "Email is not available for registration. Please use another one."
		} else if errors.As(err, &valErr) {
			fullnameError, hasFullnameErr := valErr.GetField("fullname")
			emailErr, hasEmailErr := valErr.GetField("email")
			passwordErr, hasPasswordErr := valErr.GetField("password")
			confirmPassErr, hasConfirmPassErr := valErr.GetField("confirmPassword")

			if hasFullnameErr {
				switch fullnameError {
				case service.EMPTY:
					formErrs.FullnameError = "You can't register with a blank full name."
				}
			}
			if hasEmailErr {
				switch emailErr {
				case service.EMPTY:
					formErrs.EmailError = "You can't register with a blank email."
				}
			}
			if hasPasswordErr {
				switch passwordErr {
				case service.EMPTY:
					formErrs.PasswordError = "You can't register with a blank password."
				}
			}
			if hasConfirmPassErr {
				switch confirmPassErr {
				case service.EMPTY:
					formErrs.ConfirmPasswordError = "You can't register without confirming your password."
				case service.INVALID:
					formErrs.ConfirmPasswordError = "The 2 passwords you entered don't match."
				}
			}

		} else {
			formErrs.GenericError = "An unexpected error occurred on our end. Please retry later!"
		}

		registerPage := registerpage.MakeErroredRegisterPage(&formErrs, navbar.MakeStandardNavbar(r.Context()))
		base.PageSkeleton(registerPage).Render(r.Context(), w)
		return
	}

	token, err := rc.jwtHelper.GenerateJWT(registeredUser.Email)

	if err != nil {
		registerPage := registerpage.MakeErroredRegisterPage(&formErrs, navbar.MakeStandardNavbar(r.Context()))
		base.PageSkeleton(registerPage).Render(r.Context(), w)
		return
	}

	cookie := http.Cookie{
		Name:     "authCookie",
		Value:    token,
		Path:     "/",
		MaxAge:   int(rc.jwtHelper.TokenTTL),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	w.Header().Set("HX-Redirect", "/")
}

func (rc *UserRestController) loginUser(w http.ResponseWriter, r *http.Request) {
	var userDTO service.UserLoginDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)

	formErrs := loginpage.LoginFormErrors{}
	if err != nil {
		formErrs.GenericError = "Request failed for an unknown reason"

		loginPage := loginpage.MakeErroredLoginPage(&formErrs, navbar.MakeStandardNavbar(r.Context()))

		base.PageSkeleton(loginPage).Render(r.Context(), w)
		return
	}

	loggedUser, err := rc.userService.Login(r.Context(), userDTO)

	if err != nil {
		var authErr *service.AuthError
		var entityNotFoundErr *infrastructure.EntityNotFoundError
		var valErr *service.ValidationError
		var loginPage *loginpage.ViewModel

		if errors.As(err, &authErr) {
			formErrs.GenericError = "Email or Password are incorrect"

			loginPage = loginpage.MakeErroredLoginPage(&formErrs, navbar.MakeStandardNavbar(r.Context()))
		} else if errors.As(err, &entityNotFoundErr) {
			formErrs.GenericError = "Email or Password are incorrect"

			loginPage = loginpage.MakeErroredLoginPage(&formErrs, navbar.MakeStandardNavbar(r.Context()))
		} else if errors.As(err, &valErr) {
			emailErr, hasEmailErr := valErr.GetField("email")
			passwordErr, hasPasswordErr := valErr.GetField("password")

			if hasEmailErr {
				switch emailErr {
				case service.EMPTY:
					formErrs.EmailError = "You can't log in with a blank email."
				}
			}
			if hasPasswordErr {
				switch passwordErr {
				case service.EMPTY:
					formErrs.PasswordError = "You can't log in with a blank password."
				}
			}

			loginPage = loginpage.MakeErroredLoginPage(&formErrs, navbar.MakeStandardNavbar(r.Context()))
		} else {
			formErrs.GenericError = "An unexpected error occurred on our end. Please retry later!"
			loginPage = loginpage.MakeErroredLoginPage(&formErrs, navbar.MakeStandardNavbar(r.Context()))
		}

		base.PageSkeleton(loginPage).Render(r.Context(), w)
		return
	}

	token, err := rc.jwtHelper.GenerateJWT(loggedUser.Email)

	if err != nil {
		formErrs.GenericError = "Request failed for an unknown reason"
		loginPage := loginpage.MakeErroredLoginPage(&formErrs, navbar.MakeStandardNavbar(r.Context()))
		base.PageSkeleton(loginPage).Render(r.Context(), w)
		return
	}

	cookie := http.Cookie{
		Name:     "authCookie",
		Value:    token,
		Path:     "/",
		MaxAge:   int(rc.jwtHelper.TokenTTL),
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	w.Header().Set("HX-Redirect", "/")
}

func (rc *UserRestController) deleteUser(w http.ResponseWriter, r *http.Request) {
}

func (rc *UserRestController) editUser(w http.ResponseWriter, r *http.Request) {
}
