package namaztime_test

import (
	"database/sql"
	"os"
	"reflect"
	"testing"
	"time"

	"marusya/internal/namaztime"

	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	aladhanService := NewMockAladhanService(ctrl)
	storage := NewMockDbStorage(ctrl)
	service := namaztime.NewService(aladhanService, storage, zerolog.New(os.Stderr))

	if reflect.TypeOf(service) != reflect.TypeOf(&namaztime.Service{}) {
		t.Error("Assertion error")
	}
}

func TestService_GetNamazTimeMessage(t *testing.T) {
	request := &namaztime.MarusyaRequest{
		Meta: namaztime.Meta{
			CityRu:   "Москва",
			Timezone: "Europe/Moscow",
		},
	}
	now := time.Now()
	azanTime := namaztime.AzanTime{
		ID:       1,
		City:     "Москва",
		UpdateAt: now.Add(-10 * time.Minute),
		Fajr:     now.Add(time.Hour),
		Dhuhr:    now.Add(2 * time.Hour),
		Asr:      now.Add(3 * time.Hour),
		Maghrib:  now.Add(4 * time.Hour),
		Isha:     now.Add(5 * time.Hour),
	}
	namazTextDto := &namaztime.TextDto{
		Title:       "",
		TitleRu:     "",
		TTS:         "",
		Description: "",
		Time:        "",
		TimeLeft:    "",
	}

	namazTextDto = namazTextDto.New(namaztime.Fajr, azanTime.Fajr, azanTime.Fajr.Sub(now))

	expected, err := namaztime.GetTextByTextDtoAndTemplate(namazTextDto, namaztime.MessageTemplate)
	if err != nil {
		t.Error("get expected text")
	}

	t.Run("namaz time cached", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		aladhanService := NewMockAladhanService(ctrl)
		storage := NewMockDbStorage(ctrl)
		service := namaztime.NewService(aladhanService, storage, zerolog.New(os.Stderr))
		storage.EXPECT().GetTodayAzanTimeByCity(request.Meta.CityRu).Return(&azanTime, nil)

		result, err := service.GetNamazTimeMessage(request)

		assert.Nil(t, err)
		assert.Equal(t, *expected, *result.Text)
	})

	t.Run("namaz time not cached", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		aladhanService := NewMockAladhanService(ctrl)
		storage := NewMockDbStorage(ctrl)
		service := namaztime.NewService(aladhanService, storage, zerolog.New(os.Stderr))
		storage.EXPECT().GetTodayAzanTimeByCity(request.Meta.CityRu).Return(nil, sql.ErrNoRows)
		aladhanService.EXPECT().GetTimeByCity(request.Meta.CityRu, request.Meta.Timezone).Return(&azanTime)
		storage.EXPECT().SaveAzanTime(&azanTime).Return(nil)

		result, err := service.GetNamazTimeMessage(request)

		assert.Nil(t, err)
		assert.Equal(t, *expected, *result.Text)
	})

}

func TestService_RefreshAzanTime(t *testing.T) {
	request := &namaztime.MarusyaRequest{
		Meta: namaztime.Meta{
			CityRu:   "Москва",
			Timezone: "Europe/Moscow",
		},
	}
	now := time.Now()
	azanTime := namaztime.AzanTime{
		ID:       1,
		City:     "Москва",
		UpdateAt: now.Add(-10 * time.Minute),
		Fajr:     now.Add(time.Hour),
		Dhuhr:    now.Add(2 * time.Hour),
		Asr:      now.Add(3 * time.Hour),
		Maghrib:  now.Add(4 * time.Hour),
		Isha:     now.Add(5 * time.Hour),
	}

	t.Run("refresh azan time", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		aladhanService := NewMockAladhanService(ctrl)
		storage := NewMockDbStorage(ctrl)
		service := namaztime.NewService(aladhanService, storage, zerolog.New(os.Stderr))
		aladhanService.EXPECT().GetTimeByCity(request.Meta.CityRu, request.Meta.Timezone).Return(&azanTime)
		storage.EXPECT().SaveAzanTime(&azanTime).Return(nil)

		result := service.RefreshAzanTime(request.Meta.CityRu, request.Meta.Timezone)

		assert.EqualValues(t, &azanTime, result)
	})
}
