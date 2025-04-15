package searchbar

import (
	"context"
	"io"
)

type Model struct {
	id          string
	label       string
	placeholder string
}

func Make(id, label, placeholder string) *Model {
	return &Model{
		id:          id,
		label:       label,
		placeholder: placeholder,
	}
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return component(m).Render(ctx, w)
}
