package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/amosehiguese/restaurant-api/models"
	"github.com/amosehiguese/restaurant-api/store"
	"github.com/amosehiguese/restaurant-api/types"
	"github.com/google/uuid"
)

// GetAllReservations returns all reservations
// @Summary List all reservations
// @Description Get all reservations stored in the database
// @Tags Reservation
// @Produce json
// @Router /reservations [get]
// @Success 200 {object} models.Reservation
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
func GetAllReservations(w http.ResponseWriter, r *http.Request) {
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
	result, err := q.GetAllReservations(ctx)
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

// CreateReservation writes a reservation to the database
// @Summary Creates a reservation
// @Description Creates a reservation in the database
// @Tags Reservation
// @Produce json
// @Router /reservations [post]
// @Success 200 {object} models.Reservation
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
func CreateReservation(w http.ResponseWriter, r *http.Request) {

	var payload types.ReservationPayload
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

	reservation := models.CreateReservationParams{
		TableID: payload.TableID,
		ReservationDate: parseDate(payload.ReservationDate),
		ReservationTime: parseTime(payload.ReservationTime),
		Status: models.ReservationStatusAvailable,
		CreatedAt: time.Now(),
	}

	result, err := q.CreateReservation(ctx, reservation)
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

// RetrieveReservation renders the reservation with the given id 
// @Summary Get reservation by id
// @Description RetrieveReservation returns a single reservation by id
// @Tags Reservation
// @Produce json
// @Param id path string true "reservation id"
// @Router /reservations/{id} [get]
// @Success 200 {object} models.Reservation
// @Failure 400 {object} http.StatusBadRequest
// @Failure 404 {object} http.StatusNotFound
func RetrieveReservation(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	reservationID, err := uuid.Parse(id)

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
	reservation, err := q.RetrieveReservation(ctx, reservationID)
	if err != nil {
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusNotFound,
			"msg": fmt.Sprintf("reservation with this ID %s not found", reservationID),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": reservation,
	})

}

// UpdateReservations modifies the reservation with the given id 
// @Summary Modify reservation by id
// @Description UpdateReservation modifies a single reservation by id
// @Tags Reservation
// @Produce json
// @Param id path string true "reservation id"
// @Router /reservations/{id} [patch]
// @Success 200 {object} models.Reservation
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
func UpdateReservation(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	reservationID, err := uuid.Parse(id)

	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}


	var payload types.ReservationPayload
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

	restaurant := models.UpdateReservationParams{
		ID: reservationID,
		ReservationDate: parseDate(payload.ReservationDate),
		ReservationTime: parseTime(payload.ReservationTime),
		Status: payload.Status,
	}

	q := store.GetQuery()
	err = q.UpdateReservation(ctx, restaurant)
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
	str := fmt.Sprintf("Successfully update reservation with id %s", reservationID)
	json.NewEncoder(w).Encode(resp{
		"success":true,
		"msg": str,
	})
}

// CancelReservation modifies the reservation with the given id 
// @Summary Modify reservation by id
// @Description CancelReservation modifies a single menu by id and sets its status to canceled
// @Tags Reservation
// @Produce json
// @Param id path string true "reservation id"
// @Router /reservations/{id} [patch]
// @Success 200 {object} models.Reservation
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
func DeleteReservation(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	reservationID, err := uuid.Parse(id)

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
	err = q.CancelReservation(ctx, reservationID)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	dataResp := fmt.Sprintf("Reservation with id %s is successfully deleted", reservationID)
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"msg": dataResp,
	})
}

