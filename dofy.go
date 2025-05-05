package dofy

import (
	"log"
	"time"

	"github.com/thetnaingtn/isitdayoffyet/calendarific"
)

type Holiday struct {
	Name        string
	Description string
	DaysLeft    int
}

type Dofy interface {
	GetHolidays() ([]Holiday, error)
}

type concrete struct {
	client calendarific.CalParameters
}

func NewDofy(apiKey, country, holidayType string, year int) Dofy {
	return &concrete{
		client: calendarific.CalParameters{
			ApiKey:  apiKey,
			Country: country,
			Type:    holidayType,
			Year:    int32(year),
		},
	}
}

func (c *concrete) GetHolidays() ([]Holiday, error) {
	response, err := c.client.CalData()
	if err != nil {
		log.Println("Error getting holidays: ", err)
		return nil, err
	}

	holidays := make([]Holiday, 0, len(response.Response.Holidays))
	for _, h := range response.Response.Holidays {
		now := time.Now().In(time.Local)

		duration := h.GoDate.Sub(now)
		dayLefts := int(duration.Hours() / 24)
		if dayLefts <= 0 {
			continue
		}

		holiday := Holiday{
			Name:        h.Name,
			Description: h.Description,
			DaysLeft:    dayLefts,
		}

		holidays = append(holidays, holiday)
	}

	return holidays, nil
}
