package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id" json:"productId"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description,omitempty" `
	Price       float64            `bson:"price" json:"price"`
	Quantity    uint64             `bson:"quantity" json:"quantity"`
	CreateAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type ProductRepository interface {
	CreateSingle(c context.Context, product *Product) error
	CreateMultiple(c context.Context, products []Product) error
	GetSingleById(c context.Context, productId string) (Product, error)
	GetAll(c context.Context) ([]Product, error)
	GetByStr(c context.Context, someStr string) ([]Product, error)
	UpdateSingle(c context.Context, productId string, product map[string]any) error
	UpdateMultiple(c context.Context, productIds []string, product map[string]any) error
	DeleteSingle(c context.Context, productId string) error
	DeleteMultiple(c context.Context, productIds []string) error
	DeleteAll(c context.Context) error
}
