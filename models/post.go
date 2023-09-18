package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Post representa el modelo de datos para un post
type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Titulo    string             `bson:"titulo"`
	Contenido string             `bson:"contenido"`
	Imagen    string             `bson:"imagen"`
	UsuarioID primitive.ObjectID `bson:"usuarioID"`
}
