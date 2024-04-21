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

// GetAllTables returns all tables
// @Summary List all tables
// @Description Get all tables stored in the database
// @Tags Table
// @Produce json
// @Router /tables [get]
// @Success 200 {object} models.RestaurantTable
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
func GetAllTables(w http.ResponseWriter, r *http.Request) {
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
	result, err := q.GetAllTables(ctx)
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


// CreateTable writes a table to the database
// @Summary Creates a table
// @Description Creates a table in the database
// @Tags Table
// @Produce json
// @Router /tables [post]
// @Success 200 {object} models.RestaurantTable
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
func CreateTable(w http.ResponseWriter, r *http.Request) {
	var payload types.RestaurantTable
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

	table := models.CreateTableParams{
		Number: payload.Number,
		Capacity: payload.Capacity,
		Status: models.RestaurantTableStatusAvailable,
	}

	result, err := q.CreateTable(ctx, table)
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
		"table_id": result.ID,
	})
}

// RetrieveTable renders the table with the given id 
// @Summary Get table by id
// @Description RetrieveTable returns a single table by id
// @Tags Table
// @Produce json
// @Param id path string true "table id"
// @Router /tables/{id} [get]
// @Success 200 {object} models.RestaurantTable
// @Failure 400 {object} http.StatusBadRequest
// @Failure 404 {object} http.StatusNotFound
func RetrieveTable(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	tableID, err := uuid.Parse(id)

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
	table, err := q.RetrieveTable(ctx, tableID)
	if err != nil {
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusNotFound,
			"msg": fmt.Sprintf("table with this ID %s not found", tableID),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"table": table,
	})
}

// UpdateTable modifies the table with the given id 
// @Summary Modify table by id
// @Description UpdateTable modifies a single table by id
// @Tags Table
// @Produce json
// @Param id path string true "table id"
// @Router /tables/{id} [patch]
// @Success 200 {object} models.RestaurantTable
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
func UpdateTable(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	tableID, err := uuid.Parse(id)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}

	var payload types.RestaurantTable
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


	table := &models.UpdateTableParams{
		ID: tableID,
		Number: payload.Number,
		Capacity: payload.Capacity,
		Status: payload.Status,
	}

	q := store.GetQuery()
	err = q.UpdateTable(ctx, *table)
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
	str := fmt.Sprintf("Successfully update table with id %s", tableID)
	json.NewEncoder(w).Encode(resp{
		"success":true,
		"msg": str,
	})
}

// DeleteTable removes the table with the given id 
// @Summary Removes table by id
// @Description Removes a single table by id from the database
// @Tags Table
// @Produce json
// @Param id path string true "table id"
// @Router /tables/{id} [delete]
// @Success 200 {object} string
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
func DeleteTable(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	tableID, err := uuid.Parse(id)
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

	err = q.DeleteTable(ctx, tableID)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	dataResp := fmt.Sprintf("Table with id %s is successfully deleted", id)
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"msg": dataResp,
	})
}
