package database

import (
	"fmt"
	"sign_in/config"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Conn *sqlx.DB
}

func (d *DB) InitDatabase(c *config.DB) *DB {
	var err error
	source := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Username, c.Password, c.Host, c.DBName)

	// source := fmt.Sprintf("user = %s password = %s port = %s dbname = %s sslmode=disable", c.Username, c.Password, c.Port, c.DBName)

	if d.Conn, err = sqlx.Connect("pgx", source); err != nil {
		panic(err)
	}

	d.Conn.SetConnMaxLifetime(time.Minute * 2)
	d.Conn.SetMaxIdleConns(0)
	d.Conn.SetMaxOpenConns(100)

	m, err := migrate.New("file://migrate", source)
	if err != nil {
		panic("TABLE NOT FOUND")
	}
	if err = m.Up(); err != nil {
		if err.Error() != "no change" {
			panic(err)
		}
	}
	//FOR DELETE TABLE
	// if err = m.Down(); err != nil {
	// 	if err.Error() != "no change" {
	// 		panic(err)
	// 	}
	// }
	return d
}
