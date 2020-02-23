package entities

type User struct {
	ID        string `json:"uuid,omitempty"`
	LastName  string `json:"last_name,omitempty" validate:"required"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Country   string `json:"country,omitempty" validate:"required"`
	City      string `json:"city,omitempty" validate:"required"`
	Gender    string `json:"gender,omitempty" validate:"required,oneof=Male Female"`
	BirthDate string `json:"birth_date,omitempty" validate:"required"`
}
