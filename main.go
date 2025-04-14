package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/thetnaingtn/isitdayoffyet/calendarific"
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
	fmt.Println(cfg.apiKey)
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

	fmt.Println(response.Response.Holidays[0].Name)
}
