package types

import "github.com/amosehiguese/restaurant-api/models"

type RestaurantTable struct {
	Number   int32                 `json:"number" validate:"required"`
	Capacity int32                 `json:"capacity" validate:"required"`
	Status   models.RestaurantTableStatus `json:"status" validate:"required"`
}