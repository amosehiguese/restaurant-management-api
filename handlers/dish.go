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
	} else {
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
