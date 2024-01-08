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

// GetMenu returns all menu
// @Summary List all menu
// @Description Get all menu stored in the database
// @Tags Menu
// @Produce json
// @Router /menu [get]
// @Success 200 {object} models.Menu
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
func GetMenu(w http.ResponseWriter, r *http.Request)  {
	s, e, err := paginate(w, r)
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
	result, err := q.GetAllMenu(ctx)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	
	if *e < len(result) && len(result[*s:*e]) == pageSize {
		result = result[*s:*e]
	} else if *e >= len(result) && *s < len(result) {
		result = result[*s:]
	} else if *e >= len(result) && *s >= len(result) && result != nil {
		*s = 0
		*e = pageSize
		result = result[*s:*e]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": result,
	})
}

// CreateMenu writes a menu to the database
// @Summary Creates a menu
// @Description Creates a menu in the database
// @Tags Menu
// @Produce json
// @Router /menu [post]
// @Success 200 {object} models.Menu
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
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
		"menu_id": result.ID,
	})

	
}

// RetrieveMenu renders the menu with the given id 
// @Summary Get menu by id
// @Description RetrieveMenu returns a single menu by id
// @Tags Menu
// @Produce json
// @Param id path string true "menu id"
// @Router /menu/{id} [get]
// @Success 200 {object} models.Menu
// @Failure 400 {object} http.StatusBadRequest
// @Failure 404 {object} http.StatusNotFound
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
		"menu": menu,
	})
}

// UpdateMenu modifies the menu with the given id 
// @Summary Modify menu by id
// @Description UpdateMenu modifies a single menu by id
// @Tags Menu
// @Produce json
// @Param id path string true "menu id"
// @Router /menu/{id} [patch]
// @Success 200 {object} models.Invoice
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
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

// DeleteMenu remove the menu with the given id 
// @Summary Removes menu by id
// @Description Removes a single menu by id from the database
// @Tags Menu
// @Produce json
// @Param id path string true "menu id"
// @Router /menu/{id} [delete]
// @Success 200 {object} string
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
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