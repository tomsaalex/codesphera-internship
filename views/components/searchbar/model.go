package searchbar

import (
	"context"
	"io"
)

type Model struct {
	id          string
	label       string
	placeholder string
	initialText string
}

func Make(id, label, placeholder string) *Model {
	return MakeWithValue(id, label, placeholder, "")
}

func MakeWithValue(id, label, placeholder, initialText string) *Model {
	return &Model{
		id:          id,
		label:       label,
		placeholder: placeholder,
		initialText: initialText,
	}
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return component(m).Render(ctx, w)
}
