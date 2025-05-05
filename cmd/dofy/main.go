package main

import (
	"flag"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	dofy "github.com/thetnaingtn/isitdayoffyet"
	"github.com/thetnaingtn/isitdayoffyet/ui"
)

type config struct {
	apiKey      string
	holidayType string
	country     string
	year        int
}

func main() {
	var cfg config
	flag.StringVar(&cfg.apiKey, "api-key", "", "Calendarific API Key")
	flag.StringVar(&cfg.country, "country", "th", "Country Code in iso-3166 format")
	flag.StringVar(&cfg.holidayType, "type", "national", "Holiday Type(national | local | religious | observance)")
	flag.IntVar(&cfg.year, "year", time.Now().Year(), "Year")

	flag.Parse()

	dofyClient := dofy.NewDofy(cfg.apiKey, cfg.country, cfg.holidayType, cfg.year)

	m := ui.NewAppModel(dofyClient)
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal("Error running program:", err)
	}
}
