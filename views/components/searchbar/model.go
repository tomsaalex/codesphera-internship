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

	searchBarLocation string
}

func Make(id, label, placeholder, searchBarLocation string) *Model {
	return MakeWithValue(id, label, placeholder, searchBarLocation, "")
}

func MakeWithValue(id, label, placeholder, searchBarLocation, initialText string) *Model {
	return &Model{
		id:                id,
		label:             label,
		placeholder:       placeholder,
		initialText:       initialText,
		searchBarLocation: searchBarLocation,
	}
}

func (m *Model) Render(ctx context.Context, w io.Writer) error {
	return component(m).Render(ctx, w)
}
