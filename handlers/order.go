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

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
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
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": result,
	})	
	
}
func CreateOrder(w http.ResponseWriter, r *http.Request) {
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
	id = getField(r, "dishID")
	dishID, err := uuid.Parse(id)
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
			"msg":types.ValidatorErrors(err),
		})
		return
	}

	q := store.GetQuery()
	data := models.RetrieveMenuDishParams{
		MenuID: menuID,
		ID: dishID,
	}
	dish, err := q.RetrieveMenuDish(ctx, data)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}
	
	orderData := models.CreateOrderParams{
		Status: models.OrderStatusPending,
		CreatedAt: time.Now(),

	}
	order, err := q.CreateOrder(ctx, orderData)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	orderItemData := models.AddOrderItemsParams{
		OrderID: order.ID,
		DishID: dish.ID,
		Quantity: payload.Quantity,
	}

	_, err = q.AddOrderItems(ctx, orderItemData)
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
		"order": order.ID,
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
		"order": result,
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
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": result,
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