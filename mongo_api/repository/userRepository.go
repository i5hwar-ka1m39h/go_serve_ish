package repository

import (
	"context"
	"log"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)



type userRepository struct{
	db mongo.Database
	collection string
}


func NewUserRepository(db mongo.Database, col string) model.UserRepository{
	return &userRepository{
		db : db,
		collection: col,
	}
}


func (usrRep *userRepository) CreateSingle(c context.Context, user *model.User) error{
	usrCol := usrRep.db.Collection(usrRep.collection)

	_, err := usrCol.InsertOne(c, user)

	return  err
}


func (usrRep *userRepository) CreateMultiple(c context.Context, users []model.User)error{
	usrCol := usrRep.db.Collection(usrRep.collection)

	_, err :=usrCol.InsertMany(c, users)

	return  err
}


func (usrRep *userRepository) GetSingleById(c context.Context, userId string)(model.User, error){
	usrCol := usrRep.db.Collection(usrRep.collection)


	var result model.User

	actId , err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Println("error while getting objectid", err)
		return model.User{}, err
	}
	filter := bson.D{{Key: "_id", Value: actId}}
	err = usrCol.FindOne(c, filter).Decode(&result)

	if err != nil{
		log.Println("error while finding single user by userId", err)
		return  model.User{}, err
	}

	return  result, nil

}


func (usrRep *userRepository) GetAll(c context.Context) ([]model.User, error){
	usrCol := usrRep.db.Collection(usrRep.collection)

	var result []model.User

	cursor, err := usrCol.Find(c, nil)

	if err != nil{
		log.Fatalln("error while getting all users", err)	
		return  []model.User{}, err
	}

	if err = cursor.All(c, &result); err != nil{
		log.Fatalln("error while iterating in cursor", err)
		return  []model.User{}, err
	}

	return  result, nil
}

func (usrRep *userRepository) UpdateSingle(c context.Context, userId string, user map[string]interface{})error{
	usrCol := usrRep.db.Collection(usrRep.collection)

	actId, err := primitive.ObjectIDFromHex(userId)

	if err != nil{
		log.Fatalln("error while getting objectid", err)
		return err
	}


	filter := bson.D{{Key: "_id", Value: actId}}
	update := bson.M{"$set": user}
	_, err  = usrCol.UpdateOne(c, filter, update)

	return err
}


func (usrRep *userRepository) UpdateMultiple(
	c context.Context, 
	userIds []string, 
	user map[string]interface{},
	)error{
	usrCol := usrRep.db.Collection(usrRep.collection)


	objIds := make([]primitive.ObjectID, 0, len(userIds))

	for _, id := range userIds{
		objId, err :=primitive.ObjectIDFromHex(id)

		if err != nil{
			log.Println("error while converting from hex", err)
			continue
		}

		objIds = append(objIds, objId)
	}


	filter := bson.M{
		"_id":bson.M{
			"$id":objIds,
		},
	}

	update := bson.M{
		"$set": user,
	}
	_, err := usrCol.UpdateMany(c, filter, update)

	return  err


}


func (usrRep *userRepository) DeleteSingle(c context.Context, userId string) error{
	useCol := usrRep.db.Collection(usrRep.collection)

	objId, err := primitive.ObjectIDFromHex(userId)

	if err != nil{
		return err
	}
	
	filter := bson.D{{Key: "_id", Value: objId}}
	_, err = useCol.DeleteOne(c,filter )

	return  err
}


func (usrRep *userRepository) DeleteMultiple(c context.Context, userIds []string) error{
	usrCol := usrRep.db.Collection(usrRep.collection)

	objIds := make([]primitive.ObjectID, 0, len(userIds))

	for _, id := range userIds{
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil{
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

	_, err:=usrCol.DeleteMany(c, filter)
	return  err
}

func (usrRep *userRepository) DeleteAll(c context.Context) error  {
	usrCol := usrRep.db.Collection(usrRep.collection)

	_, err:= usrCol.DeleteMany(c, bson.M{})
	return  err
}