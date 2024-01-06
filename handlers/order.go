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

// GetAllOrders returns all orders
// @Summary List all orders
// @Description Get all orders stored in the database
// @Tags Order
// @Produce json
// @Router /orders [get]
// @Success 200 {object} models.Order
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
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

// CreateOrder writes an order to the database
// @Summary Creates an order
// @Description Creates an order in the database
// @Tags Order
// @Produce json
// @Router /orders [post]
// @Success 200 {object} models.Order
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
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

// RetrieveOrder renders the order with the given id 
// @Summary Get order by id
// @Description RetrieveOrder returns a single order by id
// @Tags Order
// @Produce json
// @Param id path string true "order id"
// @Router /orders/{id} [get]
// @Success 200 {object} models.Order
// @Failure 400 {object} http.StatusBadRequest
// @Failure 404 {object} http.StatusNotFound
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

// UpdateOrder modifies the order with the given id 
// @Summary Modify order by id
// @Description UpdateOrder modifies a single order by id
// @Tags Order
// @Produce json
// @Param id path string true "order id"
// @Router /orders/{id} [patch]
// @Success 200 {object} models.Order
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
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

// DeleteOrder removes the order with the given id 
// @Summary Removes order by id
// @Description Removes a single order by id from the database
// @Tags Order
// @Produce json
// @Param id path string true "order id"
// @Router /order/{id} [delete]
// @Success 200 {object} string
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
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

// GetAllOrderItems returns all order items
// @Summary List all order items
// @Description Get all order items for a specific order stored in the database
// @Tags OrderItems
// @Produce json
// @Param id path string true "order id"
// @Router /orders/{id}/items [get]
// @Success 200 {object} models.OrderItem
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
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

// CreateOrderItem Creates an order item 
// @Summary Creates an order item
// @Description Creates an item in the database
// @Tags OrderItem
// @Produce json
// @Router /orders/{id}/items [put]
// @Success 200 {object} models.OrderItem
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
func CreateOrderItem(w http.ResponseWriter,  r *http.Request) {
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
	orderItem := models.AddOrderItemsParams {
		OrderID: orderID,
		DishID: payload.DishID,
		Quantity: payload.Quantity,
	}

	result, err := q.AddOrderItems(ctx, orderItem)
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

// UpdateOrderItem writes an order item for a specific order to the database
// @Summary Modifies an order item
// @Description UpdateOrderItem modifies an order item 
// @Tags OrderItem
// @Produce json
// @Param id path string true "order id"
// @Param itemID path string true "itemID"
// @Router /orders/{id}/{items}/{itemID} [patch]
// @Success 200 {object} models.OrderItem
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
func UpdateOrderItem(w http.ResponseWriter,  r *http.Request) {
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

// RemoveSpecificOrderItem removes an order item for a specific order.
// @Summary Removes order item by id
// @Description Removes a single order item for a specific order from the database
// @Tags OrderItem
// @Produce json
// @Param id path string true "order id"
// @Param itemID path string true "order item id"
// @Router /orders/{id}/items/{itemID} [delete]
// @Success 200 {object} string
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
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