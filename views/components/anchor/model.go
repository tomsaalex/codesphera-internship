package anchor

import (
	"context"
	"io"
)

type Model struct {
	id          string
	content     string
	destination string
}

func Make(id, content, destination string) *Model {
	return &Model{
		id:          id,
		content:     content,
		destination: destination,
	}
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return component(m).Render(ctx, w)
}
