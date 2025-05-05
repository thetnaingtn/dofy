package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	dofy "github.com/thetnaingtn/isitdayoffyet"
)

func enqueueGetHolidaysCmd() tea.Msg {
	return getHolidaysMsg{}
}

func getHolidaysCmd(client dofy.Dofy) tea.Cmd {
	return func() tea.Msg {
		holidays, err := client.GetHolidays()
		if err != nil {
			return err
		}
		return gotHolidaysMsg{holidays: holidays}
	}
}
