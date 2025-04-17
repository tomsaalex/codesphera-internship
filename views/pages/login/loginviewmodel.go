package loginpage

import (
	"context"
	"curs1_boilerplate/views/components/navbar"
	"io"
)

type LoginFormErrors struct {
	EmailError    string
	PasswordError string
	GenericError  string
}

type ViewModel struct {
	emailHasError    bool
	passwordHasError bool
	isGenericError   bool

	emailError    string
	passwordError string
	genericError  string

	emailCSSClasses    map[string]bool
	passwordCSSClasses map[string]bool

	navbar *navbar.Model
}

func makeInputClasses(hasError bool) map[string]bool {
	return map[string]bool{
		"border":        hasError,
		"border-2":      hasError,
		"border-danger": hasError,
		"form-control":  true,
	}
}

func MakeValidLoginPage(nav *navbar.Model) *ViewModel {
	formErrs := &LoginFormErrors{
		EmailError:    "",
		PasswordError: "",
		GenericError:  "",
	}

	return MakeErroredLoginPage(formErrs, nav)
}

func MakeErroredLoginPage(loginFormErrs *LoginFormErrors, navbar *navbar.Model) *ViewModel {
	emailHasError := loginFormErrs.EmailError != ""
	passwordHasError := loginFormErrs.PasswordError != ""
	isGenericError := loginFormErrs.GenericError != ""

	return &ViewModel{
		emailHasError:    emailHasError,
		passwordHasError: passwordHasError,
		isGenericError:   isGenericError,

		emailError:    loginFormErrs.EmailError,
		passwordError: loginFormErrs.PasswordError,
		genericError:  loginFormErrs.GenericError,

		emailCSSClasses:    makeInputClasses(emailHasError),
		passwordCSSClasses: makeInputClasses(passwordHasError),

		navbar: navbar,
	}
}

func (vm *ViewModel) Render(ctx context.Context, w io.Writer) error {
	return LoginPage(vm, vm.navbar).Render(ctx, w)
}
