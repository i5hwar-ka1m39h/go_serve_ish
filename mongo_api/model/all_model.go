package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// URLShortener represents a shortened URL entry in the database
type URLShortener struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OriginalURL string             `bson:"original_url" json:"original_url"`
	ShortCode   string             `bson:"short_code" json:"short_code"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	ExpireAt    *time.Time         `bson:"expire_at,omitempty" json:"expire_at,omitempty"`
	Clicks      int64              `bson:"clicks" json:"clicks"`
}

// User represents a user in the system with token mechanism
type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username       string             `bson:"username" json:"username"`
	Email          string             `bson:"email" json:"email"`
	PasswordHash   string             `bson:"password_hash" json:"password_hash"`
	AccessToken    string             `bson:"access_token,omitempty" json:"access_token,omitempty"`
	RefreshToken   string             `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	TokenExpiresAt *time.Time         `bson:"token_expires_at,omitempty" json:"token_expires_at,omitempty"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}
