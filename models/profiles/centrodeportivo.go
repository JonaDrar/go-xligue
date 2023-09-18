package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// CentroDeportivo representa el modelo de datos para un centro deportivo
type CentroDeportivo struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Nombre     string             `bson:"nombre"`
	Ciudad     string             `bson:"ciudad"`
	Direccion  string             `bson:"direccion"`
	Capacidad  int                `bson:"capacidad"`
	TarifaHora float64            `bson:"tarifaHora"`
	Deportes   []Deporte          `bson:"deportes"`
	Horarios   []Horario          `bson:"horarios"`
}
