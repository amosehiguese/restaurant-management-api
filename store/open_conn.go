package store

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/amosehiguese/restaurant-api/auth"
	"github.com/amosehiguese/restaurant-api/log"

	"github.com/amosehiguese/restaurant-api/models"
)

var (
	db *sql.DB
	err error
)

var l *log.Logger

// type Client struct {
// 	*models.Queries
// 	db *sql.DB
// 	l log.Logger
// }


func Ping() error {
	return db.Ping()
}


func SetUpDB(dbuser, dbpwd, dbhost, dbport, dbname string) {
	ctx := context.Background()
	l = log.NewLog()
	db, err = postgresConn(dbuser, dbpwd, dbhost, dbport, dbname)
	if err != nil {
		l.Infoln("unable to connect to db ->", err)
	}
	postgresMigration(dbuser, dbpwd, dbhost, dbport, dbname)

	var roles = []models.CreateRoleParams{{Name: "admin", Description: "Administrator role"}, {Name: "user", Description: "Authenticated user role"}, {Name: "anonymous", Description: "Unauthenticated user role"}}

    var user = models.CreateUserParams{FirstName: os.Getenv("ADMIN_FIRSTNAME"), LastName: os.Getenv("ADMIN_LASTNAME"), Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), PasswordHash: auth.GeneratePassword(os.Getenv("ADMIN_PASSWORD")), UserRole: 1, CreatedAt: time.Now()}	

	q := GetQuery()

	for _, r := range roles {
		q.CreateRole(ctx, r)
		l.Infof("created %v role", r.Name)
	}
	
	q.CreateUser(ctx, user)
	l.Infof("created user")
}

func GetQuery() Query {
	return models.New(db)
}
