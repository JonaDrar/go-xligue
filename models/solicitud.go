package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Solicitud representa el modelo de datos para una solicitud de partido de un equipo a otro

type Solicitud struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Solicitante primitive.ObjectID `bson:"solicitante"`
	Solicitado  primitive.ObjectID `bson:"solicitado"`
	Estado      string             `bson:"estado"`
}
