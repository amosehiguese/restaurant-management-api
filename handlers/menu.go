package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amosehiguese/restaurant-api/log"
	"github.com/amosehiguese/restaurant-api/model"
)
var l = log.NewLog()

func GetMenu(w http.ResponseWriter, r *http.Request)  {
	result, err := model.GetAllMenu()
	if err != nil {
		l.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}


	menus := NewResp("success", result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menus)
}

func CreateMenu(w http.ResponseWriter, r *http.Request)  {
	
}

func RetrieveMenu(w http.ResponseWriter, r *http.Request) {
	id := getField(r, 0)
	result, err := model.Retrieve(id)
	if err != nil {
		l.Errorln(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return 
	}

	menu := NewResp("success" ,result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&menu)
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	id := getField(r, 0)
	var m model.Menu

	err := model.Update(id, w, r, &m)
	if err != nil{
		l.Errorln(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return 
	}

	menu := NewResp("success", "Menu updated successfully")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&menu)

}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	id := getField(r, 0)
	err := model.Delete(id)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(NewError(http.StatusInternalServerError, "Internal server error"))
		return
	}

	dataResp := fmt.Sprintf("Menu with id %s is successfully deleted", id)
	resp := NewResp("success", dataResp)
	json.NewEncoder(w).Encode(&resp)
}