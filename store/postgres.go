package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)



func postgresConn(dbuser, dbpwd, dbhost, dbport, dbname string) (*sql.DB, error) {
	dbConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",dbhost, dbuser, dbpwd, dbname, dbport)
	dbX, errX := sql.Open("postgres", dbConn)
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

func postgresMigration(dbuser, dbpwd, dbhost, dbport, dbname string) error {
	// := os.Getenv("POSTGRES_URI")
	postgresURI := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",dbuser, dbpwd, dbhost, dbport, dbname )
	migSourceURL := "file://db/migrations"

	m, err := migrate.New(migSourceURL,postgresURI)
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