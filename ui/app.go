package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	dofy "github.com/thetnaingtn/isitdayoffyet"
)

type AppModel struct {
	client dofy.Dofy
	list   list.Model
}

func NewAppModel(client dofy.Dofy) AppModel {
	list := newList()

	return AppModel{client, list}
}

func (m AppModel) Init() tea.Cmd {
	return tea.Batch(enqueueGetHolidaysCmd, m.list.StartSpinner())
}

func (m AppModel) View() string {
	return m.list.View()
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case getHolidaysMsg:
		m.list.Title = "Getting holidays..."
		m.list.SetShowStatusBar(false)
		m.list.SetShowHelp(false)
		cmds = append(cmds, m.list.StartSpinner(), getHolidaysCmd(m.client))
	case gotHolidaysMsg:
		m.list.Title = "Yayy! Holidays are here 🎉🎉"
		m.list.StopSpinner()
		m.list.SetShowStatusBar(true)
		m.list.SetShowHelp(true)
		cmds = append(cmds, m.list.SetItems(holidaysToItems(msg.holidays)))
	case error:
		return m, nil
	}

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
