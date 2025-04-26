package auccreate

import (
	"context"
	"curs1_boilerplate/views/components/navbar"
	"io"
)

type AuctionCreateFormErrors struct {
	ProductNameError   string
	ProductDescError   string
	StatusError        string
	ModeError          string
	StartingPriceError string
	TargetPriceError   string
	GenericError       string
}

type ViewModel struct {
	productNameHasError   bool
	productDescHasError   bool
	statusHasError        bool
	modeHasError          bool
	startingPriceHasError bool
	targetPriceHasError   bool
	isGenericError        bool

	productNameError   string
	productDescError   string
	statusError        string
	modeError          string
	startingPriceError string
	targetPriceError   string
	genericError       string

	productNameCSSClasses   map[string]bool
	productDescCSSClasses   map[string]bool
	startingPriceCSSClasses map[string]bool
	targetPriceCSSClasses   map[string]bool

	displayTargetPrice bool
	navbar             *navbar.Model
}

func makeInputClasses(hasError bool) map[string]bool {
	return map[string]bool{
		"border":        hasError,
		"border-2":      hasError,
		"border-danger": hasError,
		"form-control":  true,
	}
}

func MakeErroredAuctionCreationPage(auctionCreateFormErrs *AuctionCreateFormErrors, displayTargetPrice bool, navbar *navbar.Model) *ViewModel {
	productNameHasError := auctionCreateFormErrs.ProductNameError != ""
	productDescHasError := auctionCreateFormErrs.ProductDescError != ""
	statusHasError := auctionCreateFormErrs.StatusError != ""
	modeHasError := auctionCreateFormErrs.ModeError != ""
	startingPriceHasError := auctionCreateFormErrs.StartingPriceError != ""
	targetPriceHasError := auctionCreateFormErrs.TargetPriceError != ""
	isGenericError := auctionCreateFormErrs.GenericError != ""

	return &ViewModel{
		productNameHasError:   productNameHasError,
		productDescHasError:   productDescHasError,
		statusHasError:        statusHasError,
		modeHasError:          modeHasError,
		startingPriceHasError: startingPriceHasError,
		targetPriceHasError:   targetPriceHasError,
		isGenericError:        isGenericError,

		productNameError:   auctionCreateFormErrs.ProductNameError,
		productDescError:   auctionCreateFormErrs.ProductDescError,
		statusError:        auctionCreateFormErrs.StatusError,
		modeError:          auctionCreateFormErrs.ModeError,
		startingPriceError: auctionCreateFormErrs.StartingPriceError,
		targetPriceError:   auctionCreateFormErrs.TargetPriceError,
		genericError:       auctionCreateFormErrs.GenericError,

		productNameCSSClasses:   makeInputClasses(productNameHasError),
		productDescCSSClasses:   makeInputClasses(productDescHasError),
		startingPriceCSSClasses: makeInputClasses(startingPriceHasError),
		targetPriceCSSClasses:   makeInputClasses(targetPriceHasError),

		displayTargetPrice: displayTargetPrice,
		navbar:             navbar,
	}
}

func MakeValidAuctionCreationPage(displayTargetPrice bool, navbar *navbar.Model) *ViewModel {
	formErrs := &AuctionCreateFormErrors{}

	return MakeErroredAuctionCreationPage(formErrs, displayTargetPrice, navbar)
}

func (vm *ViewModel) Render(ctx context.Context, w io.Writer) error {
	return AuctionCreationPage(vm, vm.navbar).Render(ctx, w)
}
