package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Liga struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Nombre      string             `bson:"nombre"`
	Pais        string             `bson:"pais"`
	FechaInicio string             `bson:"fechaInicio"`
	FechaFin    string             `bson:"fechaFin"`
	Equipos     []Equipo           `bson:"equipos"`
	Descripcion string             `bson:"descripcion"`
	Ganador     primitive.ObjectID `bson:"ganador"`
	Partidos    []Partido          `bson:"partidos"`
}
