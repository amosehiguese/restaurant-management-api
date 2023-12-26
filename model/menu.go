package model

import (
	"net/http"

	"github.com/amosehiguese/restaurant-api/log"
)

var l = log.NewLog()

type Menu struct {
	ID					int 	`bson:"_id"`
	Name				string				`json:"name" bson:"name"`
	Description			string				`json:"description" bson:"description"`
}

func GetAllMenu()([]Menu, error){
	
	return nil, nil
}

func (m *Menu) Create() (int, error) {
	

	return 0, nil
}


func Retrieve(menuId string) (*Menu, error) {
	
	return nil, nil
}

func Update(menuId string, w http.ResponseWriter, r *http.Request, m *Menu) error {
	
	return nil
}

func Delete(menuId string) error {
	

	return nil
}

