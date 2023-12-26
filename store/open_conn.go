package store

import (
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
	err error
)


func SetUpDB() {
	db, err = postgresConn()
	if err != nil {
		l.Infoln("unable to connect to db -> ", err)
	}
	postgresMigration()
}

