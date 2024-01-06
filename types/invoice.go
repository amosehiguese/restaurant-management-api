package types

import (
	"github.com/google/uuid"
)

type InvoicePayload struct {
	OrderID     uuid.UUID `json:"order_id" validate:"required"`
	TotalAmount string    `json:"total_amount" validate:"required"`
	Tax         string    `json:"tax" validate:"required"`
	Discount    string    `json:"discount"`
	GrandTotal  string    `json:"grand_total" validate:"required"`
}
type UpdateInvoicePayload struct {
	TotalAmount string    `json:"total_amount" validate:"required"`
	Tax         string    `json:"tax" validate:"required"`
	Discount    string    `json:"discount"`
	GrandTotal  string    `json:"grand_total" validate:"required"`
}