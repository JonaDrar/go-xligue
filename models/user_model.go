package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty" validate:"required"`
	Location string             `bson:"location,omitempty" json:"location,omitempty" validate:"required"`
	Title    string             `bson:"title,omitempty" json:"title,omitempty" validate:"required"`
}
