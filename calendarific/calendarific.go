package calendarific

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// The API URL to get all the holidays
const (
	CalAPI = "https://calendarific.com/api/v2/holidays?"
)

// Calendarific Parameters
type CalParameters struct {
	ApiKey   string `url:"api_key,omitempty"`
	Country  string `url:"country,omitempty"`
	Year     int32  `url:"year,omitempty"`
	Day      int32  `url:"day,omitempty"`
	Month    int32  `url:"month,omitempty"`
	Location string `url:"location,omitempty"`
	Type     string `url:"type,omitempty"`
	Language string `url:"language,omitempty"`
	Uuid     bool   `url:"uuid,omitempty"`
}

// Calendarific Response
type CalResponse struct {
	Meta struct {
		Code int `json:"code"`
	} `json:"meta"`
	Response struct {
		Holidays []Holiday `json:"holidays"`
	} `json:"response"`
}

type Holiday struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        struct {
		Iso      string `json:"iso"`
		Datetime struct {
			Year  int `json:"year"`
			Month int `json:"month"`
			Day   int `json:"day"`
		} `json:"datetime"`
	} `json:"date"`
	GoDate    time.Time   `json:"-"` // This field is not in the original response, so ignore it when parsing JSON
	Type      []string    `json:"type"`
	Locations string      `json:"locations"`
	States    interface{} `json:"states"` // sometimes its a struct, sometime its a string, so use interface
}

// We don't use this struct, since the states response json is not always a JSON
// it is sometimes a string
type States []struct {
	ID        int         `json:"id"`
	Abbrev    string      `json:"abbrev"`
	Name      string      `json:"name"`
	Exception interface{} `json:"exception"`
	Iso       string      `json:"iso"`
}

// Handle the request and fetch the data from Calendarific server
func (p *CalParameters) requestHandler() (*CalResponse, error) {

	// Initialize the response
	c := new(CalResponse)

	// Build a url query based on data passed via the parameter struct
	q, _ := query.Values(p)
	url := fmt.Sprintf("%s%s", CalAPI, q.Encode())

	// Initialize the request to the calendarific server
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c, fmt.Errorf("received error when initializing the request, error: %v", err)
	}

	// Perform request (60 seconds timeout, in case we wait for so long)
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return c, fmt.Errorf("received error when requesting the data, error: %v", err)
	}
	defer resp.Body.Close()

	// Read the content
	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return c, fmt.Errorf("received error when reading the response body, error: %v", err)
	}

	// If the status code is not 200, then we error'ed out
	if resp.StatusCode != http.StatusOK {
		return c, fmt.Errorf("received invalid status code (%v), error: %v", resp.StatusCode, err)
	}

	// Unmarshal the data
	err = json.Unmarshal(contents, &c)
	if err != nil {
		return c, fmt.Errorf("received error when unmarshalling the data, error: %v", err)
	}
	for i, h := range c.Response.Holidays {
		dt := h.Date.Datetime
		h.GoDate = time.Date(dt.Year, time.Month(dt.Month), dt.Day, 0, 0, 0, 0, time.Local)
		c.Response.Holidays[i] = h
	}

	return c, nil
}

// Request the data from Calendarific
func (p *CalParameters) CalData() (*CalResponse, error) {
	return p.requestHandler()
}
