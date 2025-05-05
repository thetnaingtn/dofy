package ui

import dofy "github.com/thetnaingtn/isitdayoffyet"

type getHolidaysMsg struct{}

type gotHolidaysMsg struct {
	holidays []dofy.Holiday
}
