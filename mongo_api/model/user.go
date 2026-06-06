package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"userId"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"`
	CreateAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// this will talk to database
type UserRepository interface {
	CreateSingle(c context.Context, user *User) error
	CreateMultiple(c context.Context, users []User) error
	GetSingleById(c context.Context, userId string) (User, error)
	GetAll(c context.Context) ([]User, error)
	GetByString(c context.Context, somestr string) ([]User, error)
	UpdateSingle(c context.Context, userId string, user map[string]any) error
	UpdateMultiple(c context.Context, userIds []string, user map[string]any) error
	DeleteSingle(c context.Context, userId string) error
	DeleteMultiple(c context.Context, userIds []string) error
	DeleteAll(c context.Context) error
}

//this will contains the business logic

type UserUsecase interface {
	CreateUser(c context.Context, user *User) error
	UpdateUser(c context.Context, userId string, user map[string]any) error
	GetUserById(c context.Context, userId string) (User, error)
	SearchUser(c context.Context, userString string) ([]User, error) //this is like find by email or something like this
}
