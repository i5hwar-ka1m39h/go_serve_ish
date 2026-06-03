package repository

import (
	"context"
	"log"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type productRepository struct {
	db         mongo.Database
	collection string
}

func NewProductRepository(db mongo.Database, col string) model.ProductRepository {
	return &productRepository{
		db:         db,
		collection: col,
	}
}

func (prodRep *productRepository) CreateSingle(c context.Context, product *model.Product) error {
	prodCol := prodRep.db.Collection(prodRep.collection)
	_, err := prodCol.InsertOne(c, product)
	return err
}

func (prodRep *productRepository) CreateMultiple(c context.Context, products []model.Product) error {
	prodCol := prodRep.db.Collection(prodRep.collection)
	_, err := prodCol.InsertMany(c, products)
	return err
}

func (prodRep *productRepository) GetSingleById(c context.Context, productId string) (model.Product, error) {
	prodCol := prodRep.db.Collection(prodRep.collection)
	var result model.Product
	actId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		log.Println("error while getting objectid", err)
		return model.Product{}, err
	}
	filter := bson.D{{Key: "_id", Value: actId}}
	err = prodCol.FindOne(c, filter).Decode(&result)
	if err != nil {
		log.Println("error while finding single product by productId", err)
		return model.Product{}, err
	}
	return result, nil
}

func (prodRep *productRepository) GetAll(c context.Context) ([]model.Product, error) {
	prodCol := prodRep.db.Collection(prodRep.collection)
	var result []model.Product
	cursor, err := prodCol.Find(c, nil)
	if err != nil {
		log.Fatalln("error while getting all products", err)
		return []model.Product{}, err
	}
	if err = cursor.All(c, &result); err != nil {
		log.Fatalln("error while iterating in cursor", err)
		return []model.Product{}, err
	}
	return result, nil
}

func (prodRep *productRepository) UpdateSingle(c context.Context, productId string, product map[string]interface{}) error {
	prodCol := prodRep.db.Collection(prodRep.collection)
	actId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		log.Fatalln("error while getting objectid", err)
		return err
	}
	filter := bson.D{{Key: "_id", Value: actId}}
	update := bson.M{"$set": product}
	_, err = prodCol.UpdateOne(c, filter, update)
	return err
}

func (prodRep *productRepository) UpdateMultiple(
	c context.Context,
	productIds []string,
	product map[string]interface{},
) error {
	prodCol := prodRep.db.Collection(prodRep.collection)
	objIds := make([]primitive.ObjectID, 0, len(productIds))
	for _, id := range productIds {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Println("error while converting from hex", err)
			continue
		}
		objIds = append(objIds, objId)
	}
	filter := bson.M{
		"_id": bson.M{
			"$in": objIds,
		},
	}
	update := bson.M{
		"$set": product,
	}
	_, err := prodCol.UpdateMany(c, filter, update)
	return err
}

func (prodRep *productRepository) DeleteSingle(c context.Context, productId string) error {
	prodCol := prodRep.db.Collection(prodRep.collection)
	objId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	_, err = prodCol.DeleteOne(c, filter)
	return err
}

func (prodRep *productRepository) DeleteMultiple(c context.Context, productIds []string) error {
	prodCol := prodRep.db.Collection(prodRep.collection)
	objIds := make([]primitive.ObjectID, 0, len(productIds))
	for _, id := range productIds {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Println("error while converting from hex", err)
			continue
		}
		objIds = append(objIds, objId)
	}
	filter := bson.M{
		"_id": bson.M{
			"$in": objIds,
		},
	}
	_, err := prodCol.DeleteMany(c, filter)
	return err
}

func (prodRep *productRepository) DeleteAll(c context.Context) error {
	prodCol := prodRep.db.Collection(prodRep.collection)
	_, err := prodCol.DeleteMany(c, bson.M{})
	return err
}
