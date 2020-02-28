package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

// User ...
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	LastName  string             `bson:"last_name" json:"last_name,omitempty" validate:"required"`
	Email     string             `bson:"email" json:"email,omitempty" validate:"required,email"`
	Country   string             `bson:"country" json:"country,omitempty" validate:"required"`
	City      string             `bson:"city" json:"city,omitempty" validate:"required"`
	Gender    string             `bson:"gender" json:"gender,omitempty" validate:"required,oneof=Male Female"`
	BirthDate CustomTime         `bson:"created_at" json:"birth_date,omitempty" validate:"required"`
}
