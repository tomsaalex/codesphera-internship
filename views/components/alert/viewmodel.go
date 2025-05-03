package custalerts

import (
	"context"
	"io"
)

type alertType string

const (
	alertDanger alertType = "alert-danger"
)

type ViewModel struct {
	alertMsg  string
	alertType alertType
}

func MakeAlertDanger(msg string) *ViewModel {
	return &ViewModel{
		alertMsg:  msg,
		alertType: alertDanger,
	}
}

func (vm *ViewModel) Render(ctx context.Context, w io.Writer) error {
	return customAlert(vm).Render(ctx, w)
}
