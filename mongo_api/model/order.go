package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string
type PaymentStatus string

const (
	StatusPending   Status = "Pending"
	StatusDeliverd  Status = "Delivered"
	StatusCancelled Status = "Cancelled"
	StatusConfirmed Status = "Confirmed"
	StatusShipped   Status = "Shipped"
)

const (
	PaymentPending   PaymentStatus = "Pending"
	PaymentPaid      PaymentStatus = "Paid"
	PaymentCancelled PaymentStatus = "Cancelled"
	PaymentFailed    PaymentStatus = "Failed"
)

type OrderItem struct {
	ProductId primitive.ObjectID `bson:"productId" json:"productId"`
	Qunatity  int64              `bson:"quantity" json:"quantity"`
	Price     float64            `bson:"quantity" json:"price"`
}

type Order struct {
	ID                primitive.ObjectID `bson:"_id" json:"orderId"`
	UserId            primitive.ObjectID `bson:"userId" json:"userId"`
	Subtotal          uint64             `bson:"subtotal" json:"subtotal"`
	Tax               uint64             `bson:"tax" json:"tax"`
	ShippingCost      uint64             `bson:"shippingCost" json:"shippingCost,omitempty"`
	Total             uint64             `bson:"total" json:"total"`
	Status            Status             `bson:"status" json:"status"`
	PaymentStatus     PaymentStatus      `bson:"paymentStatus" json:"paymentStatus"`
	ShippingAddressId primitive.ObjectID `bson:"shipAddrId" json:"shipAddrId"`

	Item []OrderItem `bson:"orderItems" json:"orderItems"`
}
