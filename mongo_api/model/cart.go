package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartProduct struct {
	ProductID primitive.ObjectID `bson:"productId" json:"productId"`
	Price     float64            `bson:"price" json:"price"`
	Quantity  uint64             `bson:"quantity" json:"quantity"`
}
type Cart struct {
	ID         primitive.ObjectID `bson:"_id" json:"cartId"`
	USerId     primitive.ObjectID `bson:"userId" json:"userId"`
	TotalPrice float64            `bson:"totalPrice" json:"totalPrice"`
	Items      []CartProduct      `bson:"items" json:"items"`
	CreateAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type CartRepository interface {
	CreateCart(c context.Context, cart *Cart) error
	UpdateCart(c context.Context, cartData map[string]any) error
	GetCart(c context.Context, cartId string) (*Cart, error)
}

type CartUsercase interface {
	CreateCart(c context.Context, cart *Cart) error
	AddToCart(c context.Context, cartData map[string]any) error
	RemoveFromCart(c context.Context, cartData map[string]any) error
	GetCartDetails(c context.Context, cartId string) (*Cart, error)
}
