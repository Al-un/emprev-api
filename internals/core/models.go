package core

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CorsConfig allows a flexible way to handle CORS stuff
type CorsConfig struct {
	Hosts   string
	Methods string
	Headers string
}

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
	IsAdmin   bool               `json:"isAdmin,omitempty" bson:"isAdmin,omitempty"`
	IsDeleted bool               `json:"isDeleted,omitempty" bson:"isDeleted"` // Soft deletion flag. Cannot omit empty value as "false" is a zero value
}

// JwtClaims extends standard claims for our User model.
//
// By including the IsAdmin and UserID fields, authorization check can be
// based on those values
type JwtClaims struct {
	IsAdmin bool   `json:"isAdmin"`
	UserID  string `json:"userId"`
	jwt.StandardClaims
}
