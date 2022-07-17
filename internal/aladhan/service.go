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

	year, _ := strconv.Atoi(date.Year)
	month := time.Month(date.Month.Number)
	day, _ := strconv.Atoi(date.Day)
	location, _ := time.LoadLocation(timezone)

	dateTime := NewDateTimeBuilder(year, month, day, location)

	azanTime := &namaztime.AzanTime{
		ID:       0,
		City:     city,
		UpdateAt: time.Now(),
		Fajr:     dateTime.prepareDateTime(response.Data.Timings.Fajr),
		Dhuhr:    dateTime.prepareDateTime(response.Data.Timings.Dhuhr),
		Asr:      dateTime.prepareDateTime(response.Data.Timings.Asr),
		Maghrib:  dateTime.prepareDateTime(response.Data.Timings.Maghrib),
		Isha:     dateTime.prepareDateTime(response.Data.Timings.Isha),
	}

	return azanTime
}

type DateTimeBuilder struct {
	Year     int
	Month    time.Month
	Day      int
	Location *time.Location
}

func NewDateTimeBuilder(year int, month time.Month, day int, location *time.Location) *DateTimeBuilder {
	return &DateTimeBuilder{
		Year:     year,
		Month:    month,
		Day:      day,
		Location: location,
	}
}

func (d *DateTimeBuilder) prepareDateTime(t string) time.Time {
	rTime := strings.Split(t, ":")
	hour, _ := strconv.Atoi(rTime[0])
	minute, _ := strconv.Atoi(rTime[1])

	return time.Date(d.Year, d.Month, d.Day, hour, minute, 1, 0, d.Location)
}
