package store

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/amosehiguese/restaurant-api/log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var l = log.NewLog()

func postgresConn() (*sqlx.DB, error) {
	postgresURI := os.Getenv("POSTGRES_URI")
	dbMaxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	dbMaxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	dbMaxLifeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	dbX, errX := sqlx.ConnectContext(context.Background(), "postgres", postgresURI)
	if errX != nil {
		return nil, errX
	}

	dbX.SetConnMaxLifetime(time.Duration(dbMaxLifeConn))
	dbX.SetMaxIdleConns(dbMaxIdleConn)
	dbX.SetMaxOpenConns(dbMaxConn)

	if errX = dbX.Ping(); errX != nil {
		defer dbX.Close()
		return nil, errX
	}

	l.Infoln("successfully connected to postgres database")

	return dbX, nil

}

func postgresMigration() error {
	postgresURI := os.Getenv("POSTGRES_URI")
	migSourceURL := "file://db/migrations"

	m, err := migrate.New(migSourceURL, postgresURI)
	if err != nil {
		l.Infoln("failed to generate migration instance ->", err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		l.Infoln("Migration up failed ->", err)
		return err
	} else if (err == migrate.ErrNoChange) {
		l.Infoln("DB is up-to-date. No migrations done ->", err)
	} else {
		l.Infoln("Migration successful")
	}

	return nil
}