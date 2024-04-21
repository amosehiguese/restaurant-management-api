package types

type DishPayload struct {
	Name        string    `json:"name" validate:"required,lte=255"`
	Description string    `json:"description" validate:"required"`
	Price       string    `json:"price" validate:"required"`
}