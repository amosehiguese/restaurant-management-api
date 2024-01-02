package types

type UserPayload struct {
	FirstName    string       `json:"first_name" validate:"required,lte=255"`
	LastName     string       `json:"last_name" validate:"required,lte=255"`
	Username     string       `json:"username" validate:"required,lte=255"`
}