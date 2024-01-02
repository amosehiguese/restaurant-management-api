package types

import "github.com/amosehiguese/restaurant-api/models"

type OrderItemPayload struct {
	Quantity int32     `json:"quantity" validate:"required"`
}

type OrderPayload struct {
	Status    models.OrderStatus  `json:"status" validate:"required"`	
}