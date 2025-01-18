package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username       string             `bson:"username" json:"username" binding:"required"`
	Password       string             `bson:"password" json:"password" binding:"required"`
	Email          string             `bson:"email" json:"email" binding:"required,email"`
	Role           string             `bson:"role" json:"role"`
	IsVerified     bool               `bson:"is_verified" json:"is_verified"`
	VerifyToken    string             `bson:"verify_token,omitempty" json:"-"`
	TokenExpiredAt time.Time          `bson:"token_expired_at,omitempty" json:"-"`
	ResetToken     string             `bson:"reset_token,omitempty" json:"-"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	LastLoginAt    time.Time          `bson:"last_login_at" json:"last_login_at"`
}

// JWT Claims
type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
