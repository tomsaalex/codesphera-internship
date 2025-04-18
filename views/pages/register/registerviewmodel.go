package registerpage

import (
	"context"
	"curs1_boilerplate/views/components/navbar"
	"io"
)

type RegisterFormErrors struct {
	FullnameError        string
	EmailError           string
	PasswordError        string
	ConfirmPasswordError string
	GenericError         string
}

type ViewModel struct {
	fullnameHasError        bool
	emailHasError           bool
	passwordHasError        bool
	confirmPasswordHasError bool
	isGenericError          bool

	fullnameError        string
	emailError           string
	passwordError        string
	confirmPasswordError string
	genericError         string

	fullnameCSSClasses        map[string]bool
	emailCSSClasses           map[string]bool
	passwordCSSClasses        map[string]bool
	confirmPasswordCSSClasses map[string]bool

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

func MakeValidRegisterPage(nav *navbar.Model) *ViewModel {
	formErrs := &RegisterFormErrors{}

	return MakeErroredRegisterPage(formErrs, nav)
}

func MakeErroredRegisterPage(registerFormErrs *RegisterFormErrors, navbar *navbar.Model) *ViewModel {
	fullnameHasError := registerFormErrs.FullnameError != ""
	emailHasError := registerFormErrs.EmailError != ""
	passwordHasError := registerFormErrs.PasswordError != ""
	confirmPasswordHasError := registerFormErrs.ConfirmPasswordError != ""
	isGenericError := registerFormErrs.GenericError != ""

	return &ViewModel{
		fullnameHasError:        fullnameHasError,
		emailHasError:           emailHasError,
		passwordHasError:        passwordHasError,
		confirmPasswordHasError: confirmPasswordHasError,
		isGenericError:          isGenericError,

		fullnameError:        registerFormErrs.FullnameError,
		emailError:           registerFormErrs.EmailError,
		passwordError:        registerFormErrs.PasswordError,
		confirmPasswordError: registerFormErrs.ConfirmPasswordError,
		genericError:         registerFormErrs.GenericError,

		fullnameCSSClasses:        makeInputClasses(fullnameHasError),
		emailCSSClasses:           makeInputClasses(emailHasError),
		passwordCSSClasses:        makeInputClasses(passwordHasError),
		confirmPasswordCSSClasses: makeInputClasses(confirmPasswordHasError),

		navbar: navbar,
	}
}

func (vm *ViewModel) Render(ctx context.Context, w io.Writer) error {
	return RegisterPage(vm, vm.navbar).Render(ctx, w)
}
