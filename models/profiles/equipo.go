package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Equipo representa el modelo de datos para un equipo
type Equipo struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	Nombre     string               `bson:"nombre"`
	Ciudad     string               `bson:"ciudad"`
	Fundacion  string               `bson:"fundacion"`
	Jugadores  []primitive.ObjectID `bson:"jugadores"`
	Entrenador string               `bson:"entrenador"`
	Estadio    string               `bson:"estadio"`
}
