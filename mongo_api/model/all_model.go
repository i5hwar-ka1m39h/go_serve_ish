package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Show represents a movie or show (example model)
type Show struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title       string   `bson:"title" json:"title"`
	Description string   `bson:"descprition" json:"desc"`
	Rating      int16    `bson:"rating" json:"rating"`
	Genre       []string `bson:"genre" json:"genre"`
}

// User represents a user in the system
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	Password  string             `bson:"password" json:"-" validate:"required,min=6"`
	FirstName string             `bson:"first_name" json:"first_name" validate:"required"`
	LastName  string             `bson:"last_name" json:"last_name" validate:"required"`
	Role      string             `bson:"role" json:"role" validate:"required,oneof=customer admin"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// Product represents a product in the catalog
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price" validate:"required,min=0"`
	CategoryID  primitive.ObjectID `bson:"category_id" json:"category_id"`
	Images      []string           `bson:"images" json:"images"`
	Stock       int                `bson:"stock" json:"stock" validate:"min=0"`
	IsActive    bool               `bson:"is_active" json:"is_active"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

// Category represents a product category
type Category struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	Description string             `bson:"description" json:"description"`
	Image       string             `bson:"image" json:"image"`
	IsActive    bool               `bson:"is_active" json:"is_active"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

// Order represents a customer order
type Order struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID          primitive.ObjectID `bson:"user_id" json:"user_id"`
	OrderItems      []OrderItem        `bson:"order_items" json:"order_items"`
	TotalAmount     float64            `bson:"total_amount" json:"total_amount"`
	Status          string             `bson:"status" json:"status" validate:"required,oneof=pending confirmed shipped delivered cancelled"`
	ShippingAddress Address            `bson:"shipping_address" json:"shipping_address"`
	BillingAddress  Address            `bson:"billing_address" json:"billing_address"`
	PaymentID       primitive.ObjectID `bson:"payment_id" json:"payment_id"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}

// OrderItem represents an item within an order
type OrderItem struct {
	ProductID   primitive.ObjectID `bson:"product_id" json:"product_id"`
	ProductName string             `bson:"product_name" json:"product_name"`
	Quantity    int                `bson:"quantity" json:"quantity" validate:"required,min=1"`
	Price       float64            `bson:"price" json:"price"`
}

// Cart represents a shopping cart
type Cart struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Items     []CartItem         `bson:"items" json:"items"`
	Total     float64            `bson:"total" json:"total"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// CartItem represents an item in the shopping cart
type CartItem struct {
	ProductID   primitive.ObjectID `bson:"product_id" json:"product_id"`
	ProductName string             `bson:"product_name" json:"product_name"`
	Quantity    int                `bson:"quantity" json:"quantity" validate:"required,min=1"`
	Price       float64            `bson:"price" json:"price"`
}

// Address represents a shipping or billing address
type Address struct {
	Street     string `bson:"street" json:"street" validate:"required"`
	City       string `bson:"city" json:"city" validate:"required"`
	State      string `bson:"state" json:"state" validate:"required"`
	Country    string `bson:"country" json:"country" validate:"required"`
	PostalCode string `bson:"postal_code" json:"postal_code" validate:"required"`
}

// Payment represents payment information
type Payment struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OrderID       primitive.ObjectID `bson:"order_id" json:"order_id"`
	UserID        primitive.ObjectID `bson:"user_id" json:"user_id"`
	Amount        float64            `bson:"amount" json:"amount"`
	PaymentMethod string             `bson:"payment_method" json:"payment_method" validate:"required,oneof=credit_card debit_card paypal bank_transfer"`
	Status        string             `bson:"status" json:"status" validate:"required,oneof=pending completed failed refunded"`
	TransactionID string             `bson:"transaction_id" json:"transaction_id"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}

// Review represents a product review
type Review struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductID primitive.ObjectID `bson:"product_id" json:"product_id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Rating    int                `bson:"rating" json:"rating" validate:"required,min=1,max=5"`
	Comment   string             `bson:"comment" json:"comment"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
