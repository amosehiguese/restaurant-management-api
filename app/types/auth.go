package types

type SignUp struct {
	FirstName	string `json:"first_name" validate:"required,lte=255"`
	LastName	string `json:"last_name" validate:"required,lte=255"`
	UserName	string `json:"username" validate:"required,lte=255"`
	Email    	string `json:"email" validate:"required,email,lte=255"`
	Password 	string `json:"password" validate:"required,lte=255"`
	UserRole 	string `json:"user_role" validate:"lte=25"`
}

type SignIn struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}


type Renew struct {
	RefreshToken	string `json:"refresh_token"`
}