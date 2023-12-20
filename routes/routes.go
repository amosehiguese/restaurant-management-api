package routes

import (
	"net/http"

	"github.com/amosehiguese/restaurant-api/handlers"
)


func HandleMenuRequest( w http.ResponseWriter, r *http.Request) {
	var err error 
	switch r.Method {
	case "GET":
		err = handlers.GetMenu(w, r)
		
	case "POST":
		err = handlers.CreateMenu(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
}
