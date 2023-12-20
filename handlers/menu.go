package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/amosehiguese/restaurant-api/db"
	"github.com/amosehiguese/restaurant-api/log"
)
var l = log.NewLog()

func GetMenu(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateMenu(w http.ResponseWriter, r *http.Request) error {
	var menu db.Menu
	err := json.NewDecoder(r.Body).Decode(&menu)
	if err != nil {
		l.Log.Error(err.Error())
		http.Error(w, "Unprocessable entity", http.StatusUnprocessableEntity)
		return err
	}

	id, err := menu.Create()
	if err != nil {
		l.Log.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return err
	}

	menu.ID = id
	menuResp := NewResp(true, menu)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menuResp)



	return nil
}