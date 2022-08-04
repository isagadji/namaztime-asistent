package namaztime

import (
	"database/sql"
	"errors"
	"time"

	"github.com/rs/zerolog"
)

const (
	Fajr    = "fajr"
	Dhuhr   = "dhuhr"
	Asr     = "asr"
	Maghrib = "maghrib"
	Isha    = "isha"
)

type Service struct {
	aladhanService AladhanService
	storage        *Storage
	logger         zerolog.Logger
}

type AladhanService interface {
	GetTimeByCity(city string, timezone string) *AzanTime
}

func NewService(aladhanService AladhanService, storage *Storage, logger zerolog.Logger) *Service {
	return &Service{
		aladhanService: aladhanService,
		storage:        storage,
		logger:         logger,
	}
}

type Msg struct {
	Text    *string
	TTSText *string
}

func (s *Service) GetNamazTimeMessage(request *MarusyaRequest) (*Msg, error) {
	azanTime, err := s.storage.getTodayAzanTimeByCity(request.Meta.CityRu)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		azanTime = s.refreshAzanTime(request.Meta.CityRu, request.Meta.Timezone)
	}

	timezone, err := time.LoadLocation(request.Meta.Timezone)
	if err != nil {
		return nil, err
	}
	s.logger.Info().Msg(timezone.String())
	now := time.Now().In(timezone)
	if now.Year() != azanTime.UpdateAt.Year() ||
		(now.Year() == azanTime.UpdateAt.Year() && now.YearDay() != azanTime.UpdateAt.YearDay()) {
		azanTime = s.refreshAzanTime(request.Meta.CityRu, request.Meta.Timezone)
	}

	var (
		actual int = 0
	)

	namazTextDto := &TextDto{
		Title:       "",
		TitleRu:     "",
		TTS:         "",
		Description: "",
		Time:        "",
		TimeLeft:    "",
	}

	azanTimes := azanTime.getAzanTimes()
	for k, a := range azanTimes {
		diff := int(a.Sub(now))
		if diff <= 0 {
			continue
		}

		if actual == 0 {
			actual = diff
			namazTextDto = namazTextDto.New(k, a, a.Sub(now))
			continue
		}

		if actual > diff {
			actual = diff
			namazTextDto = namazTextDto.New(k, a, a.Sub(now))
			continue
		}
	}

	if actual == 0 {
		fajr := azanTimes[Fajr].Add(24 * time.Hour)
		namazTextDto = namazTextDto.New(Fajr, fajr, fajr.Sub(now))
	}

	text, err := getTextByTextDtoAndTemplate(namazTextDto, namazTimeMessageTemplate)
	if err != nil {
		return nil, err
	}

	ttsText, err := getTextByTextDtoAndTemplate(namazTextDto, namazTimeMessageTemplateTTS)
	if err != nil {
		return nil, err
	}

	msg := &Msg{Text: text, TTSText: ttsText}

	return msg, nil
}

func (s *Service) refreshAzanTime(c, tz string) *AzanTime {
	azanTime := s.aladhanService.GetTimeByCity(c, tz)
	if err := s.storage.saveAzanTime(azanTime); err != nil {
		s.logger.Info().Msg(err.Error())
	}

	return azanTime
}
