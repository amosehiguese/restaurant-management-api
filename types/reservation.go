package types

import (
	"time"

	"github.com/amosehiguese/restaurant-api/models"
	"github.com/google/uuid"
)

type ReservationPayload struct {
	TableID         uuid.UUID         `json:"table_id" validate:"required"`
	ReservationDate string         `json:"reservation_date" validate:"required"`
	ReservationTime string         `json:"reservation_time" validate:"required"`
	Status          models.ReservationStatus `json:"status"`
}

type ReservationResponse struct {
	ID				uuid.UUID		`json:"id"`
	TableID         uuid.UUID         `json:"table_id" `
	ReservationDate time.Time         `json:"reservation_date" `
	ReservationTime time.Time         `json:"reservation_time" `
	Status          models.ReservationStatus `json:"status"`
	CreatedAt		time.Time		`json:"created_at"`
	UpdatedAt		time.Time		`json:"updated_at"`
}


