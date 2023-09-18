package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Horario representa el modelo de datos para un horario
type Horario struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CentroID       primitive.ObjectID `bson:"centroID"`
	DeporteID      primitive.ObjectID `bson:"deporteID"`
	Dia            string             `bson:"dia"`
	HoraInicio     string             `bson:"horaInicio"`
	HoraFin        string             `bson:"horaFin"`
	Disponibilidad int                `bson:"disponibilidad"`
}
