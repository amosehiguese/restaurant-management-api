package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amosehiguese/restaurant-api/models"
	"github.com/amosehiguese/restaurant-api/store"
	"github.com/amosehiguese/restaurant-api/types"
	"github.com/google/uuid"
)

// GetAllMenuDishes returns all dishes associated with the given menu id in the database
// @Summary List all dishes
// @Description Get all dishes associated with the given menu id stored in the database
// @Tags Dishes
// @Produce json
// @Param id path string true "menu id"
// @Router /menu/{id}/dishes [get]
// @Success 200 {object} models.Dish
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
func GetAllMenuDishes(w http.ResponseWriter, r *http.Request)  {
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

	result, err := q.GetAllMenuDishes(ctx, menuID)
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

// CreateMenuDish writes a dish to the database
// @Summary Creates a dish
// @Description Creates a dish for a given menu
// @Tags Dishes
// @Produce json
// @Param id path string true "menu id"
// @Router /menu/{id}/dishes [post]
// @Success 200 {object} models.Dish
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
func CreateMenuDish(w http.ResponseWriter, r *http.Request)  {
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

	var payload types.DishPayload
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

	if err := v.Struct(payload); err != nil {
		json.NewEncoder(w).Encode(resp{
			"error": true,
			"code": http.StatusBadRequest,
			"msg":types.ValidatorErrors(err),
		})
		return
	}

	q := store.GetQuery()
	dish := models.CreateMenuDishParams{
		Name: payload.Name,
		Description: payload.Description,
		Price: payload.Price,
		MenuID: menuID,
	}

	result, err := q.CreateMenuDish(ctx, dish)
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

// RetrieveMenuDish renders the dish with the given id 
// @Summary Get dish by id
// @Description RetrieveMenuDish returns a single dish by id
// @Tags Dish
// @Produce json
// @Param id path string true "menu id"
// @Param dishID path string true "dish id"
// @Router /menu/{id}/dishes/{dishID} [get]
// @Success 200 {object} models.Dish
// @Failure 400 {object} http.StatusBadRequest
// @Failure 404 {object} http.StatusNotFound
func RetrieveMenuDish(w http.ResponseWriter, r *http.Request)  {
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

		dishId := getField(r, "dishID")
		dishID, err := uuid.Parse(dishId)
		if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
		}

		dataIDs := models.RetrieveMenuDishParams{
			MenuID: menuID,
			ID: dishID,
		}

	q := store.GetQuery()
	result, err := q.RetrieveMenuDish(context.Background(), dataIDs)
		if err != nil {
			l.Error(err.Error())
			json.NewEncoder(w).Encode(resp{
				"success": false,
				"code": http.StatusNotFound,
				"msg": "Dish not found",
			})
			return
		}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": result,
	})
}

// UpdateMenuDish modifies the dish with the given id 
// @Summary Modify dish by id
// @Description UpdateMenuDish modifies a single dish by id
// @Tags Dish
// @Produce json
// @Param id path string true "menu id"
// @Param dishID path string true "dish id"
// @Router /menu/{id}/dishes/{dishID} [patch]
// @Success 200 {object} models.Dish
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
func UpdateMenuDish(w http.ResponseWriter, r *http.Request)  {
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

	dishId := getField(r, "dishID")
	dishID, err := uuid.Parse(dishId)
	if err != nil {
	l.Error(err.Error())
	json.NewEncoder(w).Encode(resp{
		"success": false,
		"code": http.StatusBadRequest,
		"msg": "Bad request",
	})
	return
	}

	var payload types.DishPayload
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

	q := store.GetQuery()
	dish := models.UpdateMenuDishParams{
		Name: payload.Name,
		Description: payload.Description,
		Price: payload.Price,
		MenuID: menuID,
		ID: dishID,
	}

	err = q.UpdateMenuDish(ctx, dish)
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
	str := fmt.Sprintf("Successfully update dish with id %s", dishID)
	json.NewEncoder(w).Encode(resp{
		"success":true,
		"msg": str,
	})
}

// DeleteMenuDish remove the dish with the given id 
// @Summary Removes dish by id
// @Description Removes a single dish by id under a specific menu from the database
// @Tags Dish
// @Produce json
// @Param id path string true "menu id"
// @Param dishID path string true "dish id"
// @Router /menu/{id}/dishes/{dishID} [delete]
// @Success 200 {object} string
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
func DeleteMenuDish(w http.ResponseWriter, r *http.Request)  {
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

	dishId := getField(r, "dishID")
	dishID, err := uuid.Parse(dishId)
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

	menuDish := models.DeleteMenuDishParams{
		MenuID: menuID,
		ID: dishID,
	}
	err = q.DeleteMenuDish(ctx, menuDish)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	dataResp := fmt.Sprintf("Dish with id %s is successfully deleted", dishID)
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"msg": dataResp,
	})
}
