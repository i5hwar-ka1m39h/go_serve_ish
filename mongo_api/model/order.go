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

type Order struct {
	ID                primitive.ObjectID `bson:"_id" json:"orderId"`
	UserId            string             `bson:"userId" json:"userId"`
	Subtotal          uint64             `bson:"subtotal" json:"subtotal"`
	Tax               uint64             `bson:"tax" json:"tax"`
	ShippingCost      uint64             `bson:"shippingCost" json:"shippingCost,omitempty"`
	Total             uint64             `bson:"total" json:"total"`
	Status            Status             `bson:"status" json:"status"`
	PaymentStatus     PaymentStatus      `bson:"paymentStatus" json:"paymentStatus"`
	ShippingAddressId string             `bson:"shipAddrId" json:"shipAddrId`
}
