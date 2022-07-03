package namaztime

import (
	"time"
)

var Schema = `
CREATE TABLE IF NOT EXISTS azan_time
(
    id         serial
        constraint azan_time_pk
            primary key,
    city       text,
    updated_at timestamp default current_timestamp,
    fajr       timestamp,
    dhuhr      timestamp,
    asr        timestamp,
    maghrib    timestamp,
    isha       timestamp
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
