package types

import "github.com/amosehiguese/restaurant-api/models"

type RestaurantTable struct {
	Name     string                `json:"name" validate:"required,lte=255"`
	Capacity int32                 `json:"capacity" validate:"required"`
	Status   models.RestaurantTableStatus `json:"status" validate:"required"`
}