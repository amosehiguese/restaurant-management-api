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

var l = log.NewLog()

func SetUpDB() {
	ctx := context.Background()
	db, err = postgresConn()
	if err != nil {
		l.Infoln("unable to connect to db ->", err)
	}
	postgresMigration()

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

func GetQuery() *models.Queries {
	return models.New(db)
}
