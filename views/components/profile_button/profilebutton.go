package profilebutton

import (
	"context"
	"io"
)

type Model struct {
	userEmail string
}

func Make(userEmail string) *Model {
	return &Model{
		userEmail: userEmail,
	}
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return component(m).Render(ctx, w)
}
