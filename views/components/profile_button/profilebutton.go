package profilebutton

import (
	"context"
	"io"
)

type Model struct {
}

func Make() *Model {
	return &Model{}
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return component().Render(ctx, w)
}
