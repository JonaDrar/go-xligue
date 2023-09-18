package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Jugador representa el modelo de datos para un jugador
type Jugador struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Nombre            string             `bson:"nombre"`
	Posicion          string             `bson:"posicion"`
	FotoPerfil        string             `bson:"fotoPerfil"`
	EquipoID          primitive.ObjectID `bson:"equipoID"`
	PartidosJugados   int                `bson:"partidosJugados"`
	GolesMarcados     int                `bson:"golesMarcados"`
	Asistencias       int                `bson:"asistencias"`
	TarjetasAmarillas int                `bson:"tarjetasAmarillas"`
	TarjetasRojas     int                `bson:"tarjetasRojas"`
	FechaNacimiento   string             `bson:"fechaNacimiento"`
	Nacionalidad      string             `bson:"nacionalidad"`
}
