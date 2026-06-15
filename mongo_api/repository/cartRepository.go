package repository

import (
	"context"
	"log"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type cartRepository struct {
	database mongo.Database
	coll     string
}

func NewCartRepository(db mongo.Database, collection string) model.CartRepository {
	return &cartRepository{
		database: db,
		coll:     collection,
	}
}

func (cartRep *cartRepository) CreateCart(c context.Context, cart *model.Cart) error {
	coll := cartRep.database.Collection(cartRep.coll)

	_, err := coll.InsertOne(c, cart )

	return err
}
func (cartRep *cartRepository) UpdateCart(c context.Context, cartId string, cartData map[string]any) error {
	col := cartRep.database.Collection(cartRep.coll)


	objId , err:= primitive.ObjectIDFromHex(cartId)
	if err != nil{
		log.Println("error converting objectid from ", err)
		return  err
	}
	_, err =col.UpdateByID(c, objId, cartData)
	return err
	
}
func (cartRep *cartRepository) GetCart(c context.Context, cartId string) (*model.Cart, error){
	col := cartRep.database.Collection(cartRep.coll)

	var cart model.Cart
	objId , err := primitive.ObjectIDFromHex(cartId)
	if err != nil{
		log.Println("error converting objectid from ", err)
		return  &cart, err
	}

	filter := bson.D{{Key: "_id", Value: objId}}
	err = col.FindOne(c, filter).Decode(&cart)
	if err != nil{
		log.Println("error getting single from db", err)
		return  &cart, err
	}

	return  &cart, nil
}
