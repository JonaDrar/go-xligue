package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Deporte representa el modelo de datos para un deporte
type Deporte struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Nombre      string             `bson:"nombre"`
	Descripcion string             `bson:"descripcion"`
	Reglas      string             `bson:"reglas"`
	Imagen      string             `bson:"imagen"`
}
