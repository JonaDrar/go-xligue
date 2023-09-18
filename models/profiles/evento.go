package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Evento struct {
	Tipo    string             `bson:"tipo"` // Puede ser "gol", "asistencia", "tarjeta"
	Jugador primitive.ObjectID `bson:"jugador"`
	Minuto  int                `bson:"minuto"`
}
