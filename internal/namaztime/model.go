package namaztime

import (
	"time"
)

var DDL = `
CREATE TABLE IF NOT EXISTS azan_time
(
    id         serial
        constraint azan_time_pk
            primary key,
    city       text,
    updated_at timestamp with time zone default current_timestamp,
    fajr       timestamp with time zone,
    dhuhr      timestamp with time zone,
    asr        timestamp with time zone,
    maghrib    timestamp with time zone,
    isha       timestamp with time zone
);

alter table azan_time
    owner to postgres;

create unique index IF NOT EXISTS azan_time_city_uindex
    on azan_time (city);
`

type AzanTime struct {
	ID       int64     `db:"id"`
	City     string    `db:"city"`
	UpdateAt time.Time `db:"updated_at"`
	Fajr     time.Time `db:"fajr"`
	Dhuhr    time.Time `db:"dhuhr"`
	Asr      time.Time `db:"asr"`
	Maghrib  time.Time `db:"maghrib"`
	Isha     time.Time `db:"isha"`
}

func (t AzanTime) getAzanTimes() map[string]time.Time {
	return map[string]time.Time{
		Fajr:    t.Fajr,
		Dhuhr:   t.Dhuhr,
		Asr:     t.Asr,
		Maghrib: t.Maghrib,
		Isha:    t.Isha,
	}
}
