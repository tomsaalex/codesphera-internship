package button

import (
	"context"
	"io"
)

type Model struct {
	id        string
	content   string
	getAddr   string
	getTarget string
}

func Make(id, content, getAddr, getTarget string) *Model {
	return &Model{
		id:        id,
		content:   content,
		getAddr:   getAddr,
		getTarget: getTarget,
	}
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return component(m).Render(ctx, w)
}
