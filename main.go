package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thetnaingtn/isitdayoffyet/calendarific"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)
)

type config struct {
	apiKey      string
	holidayType string
	country     string
	year        int
}

type model struct {
	list list.Model
}

type holiday struct {
	name        string
	description string
	daysLeft    int
}

func (h holiday) Title() string {
	return fmt.Sprintf(`%d day%s left until %s`, h.daysLeft, mayBePlural(h.daysLeft), h.name)
}
func (h holiday) Description() string { return h.description }
func (h holiday) FilterValue() string { return h.name }

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	return appStyle.Render(m.list.View())
}

func newModel(holidays []holiday) *model {
	items := make([]list.Item, 0, len(holidays))
	for _, holiday := range holidays {
		items = append(items, holiday)
	}

	delegateItems := list.NewDefaultDelegate()
	holidaysList := list.New(items, delegateItems, 0, 0)
	holidaysList.Title = "🎉🎉Holidays🎉🎉"
	holidaysList.Styles.Title = titleStyle

	m := model{
		list: holidaysList,
	}

	return &m
}

func mayBePlural(d int) string {
	if d > 1 {
		return "s"
	}

	return ""
}

func main() {
	var cfg config
	flag.StringVar(&cfg.apiKey, "api-key", "", "Calendarific API Key")
	flag.StringVar(&cfg.country, "country", "th", "Country Code in iso-3166 format")
	flag.StringVar(&cfg.holidayType, "type", "national", "Holiday Type(national | local | religious | observance)")
	flag.IntVar(&cfg.year, "year", time.Now().Year(), "Year")

	flag.Parse()
	parameters := calendarific.CalParameters{
		ApiKey:  cfg.apiKey,
		Country: cfg.country,
		Type:    cfg.holidayType,
		Year:    int32(cfg.year),
	}

	response, err := parameters.CalData()
	if err != nil {
		log.Fatal(err)
	}

	holidays := make([]holiday, 0, len(response.Response.Holidays))
	for _, h := range response.Response.Holidays {
		now := time.Now().In(time.Local)

		duration := h.GoDate.Sub(now)
		dayLefts := int(duration.Hours() / 24)
		if dayLefts <= 0 {
			continue
		}

		holiday := holiday{
			name:        h.Name,
			description: h.Description,
			daysLeft:    dayLefts,
		}

		holidays = append(holidays, holiday)
	}

	if _, err := tea.NewProgram(newModel(holidays), tea.WithAltScreen()).Run(); err != nil {
		log.Fatal("Error running program:", err)
	}
}
