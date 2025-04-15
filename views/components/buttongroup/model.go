package buttongroup

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type Model struct {
	id      string
	label   string
	buttons []templ.Component
}

func Make(id, label string, buttons []templ.Component) *Model {
	return &Model{
		id:      id,
		label:   label,
		buttons: buttons,
	}
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return component(m).Render(ctx, w)
}
