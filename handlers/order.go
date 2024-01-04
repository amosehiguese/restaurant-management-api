package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/amosehiguese/restaurant-api/models"
	"github.com/amosehiguese/restaurant-api/store"
	"github.com/amosehiguese/restaurant-api/types"
	"github.com/google/uuid"
)

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
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
	result, err := q.GetAllOrders(ctx)
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

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var payload types.OrderPayload 
	
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
	order := models.CreateOrderParams {
		Status: payload.Status,
		CreatedAt: time.Now(),
	}

	result, err := q.CreateOrder(ctx, order)
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


func RetrieveOrder(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	orderID, err := uuid.Parse(id)
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
	result, err := q.RetrieveOrder(ctx, orderID)
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
		"order": map[string]any{
			"order_id": result.ID,
			"status": result.Status,
			"created_at": result.CreatedAt,
			"updated_at": result.UpdatedAt.Time,
		},
	})
}
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	orderID, err := uuid.Parse(id)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}

	var payload types.OrderPayload
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
			"msg": types.ValidatorErrors(err),
		})
		return
	}

	q := store.GetQuery()
	order := models.UpdateOrderParams{
		ID: orderID,
		Status: payload.Status,
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true },
	}

	err = q.UpdateOrder(ctx, order)
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
	str := fmt.Sprintf("Successfully update order with id %s", orderID)
	json.NewEncoder(w).Encode(resp{
		"success":true,
		"msg": str,
	})
	
}
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	orderID, err := uuid.Parse(id)
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
	err = q.DeleteOrder(ctx, orderID)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}
	
	dataResp := fmt.Sprintf("Order with id %s is successfully deleted", orderID)
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"msg": dataResp,
	})	
}

func GetAllOrderItems(w http.ResponseWriter,  r *http.Request) {
	id := getField(r, "id")
	orderID, err := uuid.Parse(id)
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
	result, err := q.GetAllOrderItems(ctx, orderID)
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

// func CreateOrderItem(w http.ResponseWriter,  r *http.Request) {
// 	id := getField(r, "id")
// 	orderID, err := uuid.Parse(id)
// 	if err != nil {
// 		l.Error(err.Error())
// 		json.NewEncoder(w).Encode(resp{
// 			"success": false,
// 			"code": http.StatusBadRequest,
// 			"msg": "Bad request",
// 		})
// 		return
// 	}
// 	var payload types.CreateOrderItemPayload 
	
// 	err = json.NewDecoder(r.Body).Decode(&payload)
// 	if err != nil {
// 		l.Errorln(err)
// 		json.NewEncoder(w).Encode(resp{
// 			"success": false,
// 			"code": http.StatusUnprocessableEntity,
// 			"msg": "Unprocessable entity",
// 		})
// 		return
// 	}

// 	v := types.NewValidator()

// 	if err := v.Struct(payload); err != nil {
// 		json.NewEncoder(w).Encode(resp{
// 			"error": true,
// 			"code": http.StatusBadRequest,
// 			"msg":types.ValidatorErrors(err),
// 		})
// 		return
// 	}

// 	q := store.GetQuery()
// 	orderItem := models.AddOrderItemsParams {
// 		OrderID: orderID,
// 		DishID: payload.DishID,
// 		Quantity: payload.Quantity,
// 	}

// 	result, err := q.AddOrderItems(ctx, orderItem)
// 		if err != nil {
// 		l.Error(err.Error())
// 		json.NewEncoder(w).Encode(resp{
// 			"success": false,
// 			"code": http.StatusInternalServerError,
// 			"msg": "Internal server error",
// 		})
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(resp{
// 		"success": true,
// 		"data": result.ID,
// 	})

// }

func CreateOrUpdateOrderItem(w http.ResponseWriter,  r *http.Request) {
	id := getField(r, "itemID")
	orderItemID, err := uuid.Parse(id)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}

	var payload types.OrderItemPayload
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
			"msg": types.ValidatorErrors(err),
		})
		return
	}

	q := store.GetQuery()
	orderItem := models.UpdateOrderItemParams{
		ID: orderItemID,
		Quantity: payload.Quantity,
		DishID: payload.DishID,
	}

	err = q.UpdateOrderItem(ctx, orderItem)
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
	str := fmt.Sprintf("Successfully update orderItem with id %s", orderItemID)
	json.NewEncoder(w).Encode(resp{
		"success":true,
		"msg": str,
	})
}


func RemoveSpecificOrderItem(w http.ResponseWriter,  r *http.Request) {
	id := getField(r, "id")
	orderID, err := uuid.Parse(id)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}

	id = getField(r, "itemID")
	orderItemID, err := uuid.Parse(id)
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
	dataOrderItem := models.RemoveSpecificOrderItemParams{
		ID: orderItemID,
		OrderID: orderID,
	}
	err = q.RemoveSpecificOrderItem(ctx, dataOrderItem)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}
	
	dataResp := fmt.Sprintf("OrderItem with id %s is successfully deleted", orderItemID)
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"msg": dataResp,
	})
}