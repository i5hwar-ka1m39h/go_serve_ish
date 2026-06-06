package repository

import (
	"context"
	"log"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type orderRepository struct {
	db         mongo.Database
	collection string
}

func NewOrderRepository(db mongo.Database, col string) model.OrderRepository {
	return &orderRepository{
		db:         db,
		collection: col,
	}
}

func (ordRep *orderRepository) CreateSingle(c context.Context, order *model.Order) error {
	ordCol := ordRep.db.Collection(ordRep.collection)

	_, err := ordCol.InsertOne(c, order)
	return err
}

func (ordRep *orderRepository) GetSingleId(c context.Context, orderId string) (model.Order, error) {
	ordCol := ordRep.db.Collection(ordRep.collection)

	objId, err := primitive.ObjectIDFromHex(orderId)

	if err != nil {
		log.Println("failed to convert string to objid", err)
		return model.Order{}, err
	}

	var order model.Order
	filter := bson.D{{Key: "_id", Value: objId}}
	err = ordCol.FindOne(c, filter).Decode(&order)
	if err != nil {
		log.Println("failed to get order", err)
		return model.Order{}, err
	}

	return order, nil
}

func (ordRep *orderRepository) GetAllForUser(c context.Context, userId string) ([]model.Order, error) {
	ordCol := ordRep.db.Collection(ordRep.collection)
	objId, err := primitive.ObjectIDFromHex(userId)

	var orders []model.Order
	if err != nil {
		log.Println("error converting string to obj id", err)
		return orders, err
	}

	filter := bson.D{{Key: "userId", Value: objId}}

	cursor, err := ordCol.Find(c, filter)

	if err != nil {
		log.Println("error occred while getting all the orders", err)
		return orders, err
	}

	if err = cursor.All(c, &orders); err != nil {
		log.Println("failed to get all from cursor")
		return orders, err
	}

	return orders, nil

}

func (ordRep *orderRepository) UpdateSingle(c context.Context, orderId string, order model.Order) error {
	ordCol := ordRep.db.Collection(ordRep.collection)

	objId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		log.Println("error converting objectId from hex", err)
		return err
	}

	_, err = ordCol.UpdateByID(c, objId, order)
	if err != nil {
		log.Println("error occured while updating the order", err)
		return err
	}

	return nil

}
