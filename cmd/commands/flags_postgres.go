package commands

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresFlags struct {
	Host     string `kong:"required,group=Postgres,name=postgres-host,env=POSTGRES_HOST"`
	Db       string `kong:"required,group=Postgres,name=postgres-db,env=POSTGRES_DB"`
	User     string `kong:"required,group=Postgres,name=postgres-user,env=POSTGRES_USER"`
	Password string `kong:"required,group=Postgres,name=postgres-password,env=POSTGRES_PASSWORD"`
	Port     int    `kong:"required,group=Postgres,name=postgres-port,env=POSTGRES_PORT"`
}

func (f PostgresFlags) Init() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", f.buildDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (f PostgresFlags) buildDSN() string {
	return fmt.Sprintf(
		"host='%s' user='%s' password='%s' dbname='%s' port=%d sslmode=disable TimeZone=Europe/Moscow",
		f.Host,
		f.User,
		f.Password,
		f.Db,
		f.Port,
	)
}
