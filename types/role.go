package types

type RolePayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`	
}
