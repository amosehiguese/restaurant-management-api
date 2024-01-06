package types

import (
	"github.com/amosehiguese/restaurant-api/models"
	"github.com/google/uuid"
)

type ReservationPayload struct {
	TableID         uuid.UUID         `json:"table_id" validate:"required"`
	ReservationDate string         `json:"reservation_date" validate:"required"`
	ReservationTime string         `json:"reservation_time" validate:"required"`
	Status          models.ReservationStatus `json:"status"`
}


