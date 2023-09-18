package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Usuario representa el modelo de datos para un usuario
type Usuario struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Nombre   string             `bson:"nombre"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Rol      string             `bson:"rol"` // Puede ser "jugador", "entrenador", "administrador", por defecto es hincha.
}
