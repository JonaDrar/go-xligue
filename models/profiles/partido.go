package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Partido struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Fecha           string             `bson:"fecha"`
	Hora            string             `bson:"hora"`
	EquipoLocal     primitive.ObjectID `bson:"equipoLocal"`
	EquipoVisitante primitive.ObjectID `bson:"equipoVisitante"`
	LigaID          primitive.ObjectID `bson:"ligaID"`
	Resultado       string             `bson:"resultado"`
	Arbitro         string             `bson:"arbitro"`
	Eventos         []Evento           `bson:"eventos"`
}
