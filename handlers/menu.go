package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amosehiguese/restaurant-api/models"
	"github.com/amosehiguese/restaurant-api/store"
	"github.com/amosehiguese/restaurant-api/types"
	"github.com/google/uuid"
)

func GetMenu(w http.ResponseWriter, r *http.Request)  {
	q := store.GetQuery()
	menus, err := q.GetAllMenu(ctx)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": menus,
	})
}

func CreateMenu(w http.ResponseWriter, r *http.Request)  {
	var payload types.MenuPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		l.Errorln(err)
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusUnprocessableEntity,
			"msg": "Unprocessable entity",
		})
		return
	}

	v := types.NewValidator()

	if err := v.Struct(payload); err != nil {
		json.NewEncoder(w).Encode(resp{
			"error": true,
			"code": http.StatusBadRequest,
			"msg":types.ValidatorErrors(err),
		})
		return
	}

	q := store.GetQuery()

	menu := models.CreateMenuParams{
		Name: payload.Name,
		Description: payload.Description,
	}
	result, err := q.CreateMenu(ctx, menu)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": result.ID,
	})

	
}

func RetrieveMenu(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	menuID, err := uuid.Parse(id)

	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}
	q := store.GetQuery()
	menu, err := q.RetrieveMenu(ctx, menuID)
	if err != nil {
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusNotFound,
			"msg": fmt.Sprintf("menu with this ID %s not found", menuID),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": menu,
	})
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	menuID, err := uuid.Parse(id)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}

	var payload types.MenuPayload
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		l.Errorln(err)
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusUnprocessableEntity,
			"msg": "Unprocessable entity",
		})
		return
	}

	v := types.NewValidator()

	if err := v.Struct(payload); err != nil{
		json.NewEncoder(w).Encode(resp{
			"error": true,
			"code": http.StatusBadRequest,
			"msg":types.ValidatorErrors(err),
		})
		return
	}


	menu := &models.UpdateMenuParams{
		ID: menuID,
		Name: payload.Name,
		Description: payload.Description,
	}

	q := store.GetQuery()
	err = q.UpdateMenu(ctx, *menu)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	str := fmt.Sprintf("Successfully update menu with id %s", menuID)
	json.NewEncoder(w).Encode(resp{
		"success":true,
		"msg": str,
	})

}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	menuID, err := uuid.Parse(id)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}
	q := store.GetQuery()

	err = q.DeleteMenu(ctx, menuID)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	dataResp := fmt.Sprintf("Menu with id %s is successfully deleted", id)
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"msg": dataResp,
	})
}