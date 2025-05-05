package ui

import "github.com/charmbracelet/bubbles/list"

func newList() list.Model {

	defaultDelegate := list.NewDefaultDelegate()

	listModel := list.New([]list.Item{}, defaultDelegate, 0, 0)

	return listModel
}
