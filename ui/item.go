package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	dofy "github.com/thetnaingtn/isitdayoffyet"
)

type item struct {
	holiday dofy.Holiday
}

func (i item) Title() string {
	return fmt.Sprintf(`%d day%s left until %s`, i.holiday.DaysLeft, mayBePlural(i.holiday.DaysLeft), i.holiday.Name)
}
func (i item) Description() string { return i.holiday.Description }
func (i item) FilterValue() string { return i.holiday.Name }

func holidaysToItems(h []dofy.Holiday) []list.Item {
	items := make([]list.Item, 0, len(h))
	for _, holiday := range h {
		items = append(items, item{holiday: holiday})
	}

	return items
}

func mayBePlural(d int) string {
	if d > 1 {
		return "s"
	}

	return ""
}
