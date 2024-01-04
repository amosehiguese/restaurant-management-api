package types

import (
	"github.com/amosehiguese/restaurant-api/models"
	"github.com/google/uuid"
)

type OrderItemPayload struct {
	Quantity int32     `json:"quantity" validate:"required"`
	DishID   uuid.UUID `json:"dish_id" validate:"required"`
}

type CreateOrderItemPayload struct {
	Quantity int32     `json:"quantity" validate:"required"`
}

type OrderPayload struct {
	Status    models.OrderStatus  `json:"status" validate:"required"`	
}

