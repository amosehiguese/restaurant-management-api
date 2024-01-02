package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)



func postgresConn() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")


	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",host, username, password, dbname, port)



	dbX, errX := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errX
	}

	if err := dbX.Ping(); err != nil {
		defer dbX.Close()
		return nil, err
	}

	log.Println("successfully connected to postgres database")

	return dbX, nil
}

func postgresMigration() error {
	postgresURI := os.Getenv("POSTGRES_URI")
	migSourceURL := "file://db/migrations"

	m, err := migrate.New(migSourceURL, postgresURI)
	if err != nil {
		log.Println("failed to generate migration instance ->", err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Println("Migration up failed ->", err)
		return err
	} else if (err == migrate.ErrNoChange) {
		log.Println("DB is up-to-date. No migrations done ->",err)
	} else{
		log.Println("Migration successful")
	}

	return nil
}