package menu

import (
	"github.com/amosehiguese/restaurant-api/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var menuCollection = store.GetCollection("menu")

type menu struct {
	ID					primitive.ObjectID
	Name				string
	Description			string
}

// func Get()([]menu, error){
// 	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
// 	defer cancel()

// 	var menu []menu
// 	cur, err := menuCollection.Find(ctx, bson.D{{}})
// 	if err != nil {
// 		log.Prin
// 	}
// }