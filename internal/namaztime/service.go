package namaztime

import (
	"database/sql"
	"errors"
	"time"

	"github.com/rs/zerolog"
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

func (s Service) GetNamazTimeFromCity(request *MarusyaRequest) (*string, error) {
	azanTime, err := s.storage.getAzanTimeByCity(request.Meta.CityRu)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if azanTime == nil {
		if err = s.creatAndUpdate(request, azanTime); err != nil {
			return nil, err
		}
	}

	if (time.Now().Year() == azanTime.UpdateAt.Year() && time.Now().YearDay() != azanTime.UpdateAt.YearDay()) ||
		time.Now().Year() != azanTime.UpdateAt.Year() {
		if err = s.creatAndUpdate(request, azanTime); err != nil {
			return nil, err
		}
	}

	s.logger.Info().Msg(azanTime.City)

	timezone, err := time.LoadLocation(request.Meta.Timezone)
	if err != nil {
		return nil, err
	}
	now := time.Now().In(timezone)
	diff := now.Sub(azanTime.Fajr)

	s.logger.Info().Msg(diff.String())

	//tmpl, err := template.New("namaz-time").Parse(namazTimeResponseTextTemplate)
	//if err != nil {
	//	return nil, err
	//}
	//
	//resultBuffer := &bytes.Buffer{}
	//if err = tmpl.Execute(resultBuffer, namazTime); err != nil {
	//	return nil, err
	//}
	//
	//result := resultBuffer.String()

	//return &result, nil
	return nil, nil
}

func (s Service) creatAndUpdate(request *MarusyaRequest, azanTime *AzanTime) error {
	azanTime = s.aladhanService.GetTimeByCity(request.Meta.CityRu, request.Meta.Timezone)

	return s.storage.saveAzanTime(azanTime)
}

func checkTime(aTime time.Time) {

}
