package aladhan

import (
	"strconv"
	"strings"
	"time"

	"marusya/internal/namaztime"
)

type Service struct {
	Client *Client
}

func NewService(client *Client) *Service {
	return &Service{Client: client}
}

func (s *Service) GetTimeByCity(city string, timezone string) *namaztime.AzanTime {
	response, err := s.Client.GetTimeByCity(city)
	if err != nil {
		return nil
	}
	date := response.Data.Date.Gregorian
	timings := response.Data.Timings

	year, _ := strconv.Atoi(date.Year)
	month := time.Month(date.Month.Number)
	day, _ := strconv.Atoi(date.Day)
	location, _ := time.LoadLocation(timezone)

	return &namaztime.AzanTime{
		City:     city,
		UpdateAt: time.Now(),
		Fajr:     prepareDateTime(year, month, day, timings.Fajr, location),
		Dhuhr:    prepareDateTime(year, month, day, timings.Dhuhr, location),
		Asr:      prepareDateTime(year, month, day, timings.Asr, location),
		Maghrib:  prepareDateTime(year, month, day, timings.Maghrib, location),
		Isha:     prepareDateTime(year, month, day, timings.Isha, location),
	}
}

func prepareDateTime(year int, month time.Month, day int, t string, location *time.Location) time.Time {
	rTime := strings.Split(t, ":")
	hour, _ := strconv.Atoi(rTime[0])
	minute, _ := strconv.Atoi(rTime[1])

	return time.Date(year, month, day, hour, minute, 1, 0, location)
}
