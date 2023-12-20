package db

import (
	"context"
	"time"

	"github.com/amosehiguese/restaurant-api/log"
	"github.com/amosehiguese/restaurant-api/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var menuCollection = store.GetCollection("menu")
var l = log.NewLog()

type Menu struct {
	ID					primitive.ObjectID
	Name				string
	Description			string
}

func GetAllMenu()([]Menu, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var menus []Menu
	cur, err := menuCollection.Find(ctx, bson.D{{}})
	if err != nil {
		l.Log.Error(err.Error())
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var m Menu
		if err = cur.Decode(&m); err != nil {
			l.Log.Sugar().Infoln("Unable to decode ->", err)
			return nil,err
		}

		menus = append(menus, m)
	}

	if err := cur.Err(); err != nil {
		l.Log.Info(err.Error())
		return nil, err
	}

	cur.Close(context.TODO())
	return menus, nil
}

func (m *Menu) Create() (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	result, err := menuCollection.InsertOne(ctx, m)
	if err  != nil {
		l.Log.Sugar().Infoln("Insert not successful ->", err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
