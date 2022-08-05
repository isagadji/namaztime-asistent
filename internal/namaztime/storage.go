package namaztime

import (
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	db.MustExec(DDL)
	return &Storage{db: db}
}

func (s *Storage) GetTodayAzanTimeByCity(city string) (*AzanTime, error) {
	azanTime := &AzanTime{}
	err := s.db.Get(azanTime, "SELECT * FROM azan_time WHERE city=$1 LIMIT 1", city)

	return azanTime, err
}

func (s *Storage) SaveAzanTime(azanTime *AzanTime) error {
	queryString := "INSERT INTO azan_time (city, fajr, dhuhr, asr, maghrib, isha) " +
		"VALUES (:city, :fajr, :dhuhr, :asr, :maghrib, :isha) " +
		"ON CONFLICT (city) DO UPDATE " +
		"SET updated_at=now(), fajr=:fajr, dhuhr=:dhuhr, asr=:asr, maghrib=:maghrib, isha=:isha;"
	_, err := s.db.NamedExec(queryString, azanTime)
	if err != nil {
		return err
	}

	return nil
}
