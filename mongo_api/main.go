package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Collection {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("error occured while loadong env", err)
	}

	connection_url := os.Getenv("MONGO_URI")

	conn_options := options.Client().ApplyURI(connection_url)

	mongo_client, err := mongo.Connect(context.Background(), conn_options)

	if err != nil {
		log.Println("error occured while creating mongo client ", err)
	}

	err = mongo_client.Ping(context.Background(), nil)
	if err != nil {
		log.Println("unable to ping the ")
	} else {
		fmt.Println("connected to mongodb")
	}

	collection := mongo_client.Database("e-commerse").Collection("test_collection")

	return collection
}

func insertData(ctx context.Context, collection *mongo.Collection) error {
	select{
	case <- ctx.Done():
		return  ctx.Err()
	default:

	}	





	my_movie := Show{
		Title:       "land hamar chhota",
		Description: "A action horror pic expressing monsterization of small dicks",
		Rating:      4,
		Genre:       []string{"action", "horror", "psycological", "gore", "melodrama"},
	}
	result , err := collection.InsertOne(ctx, my_movie)

	if err != nil{
		log.Println("error occured while inserting the document", err)
	
		return fmt.Errorf("erro inside the insertData %w", err)
	}

	fmt.Println("inserted the document ", result)
	return  nil

}

type Show struct {
	ID          string   `bson:"_id,omitempty" json:"id"`
	Title       string   `bson:"title" json:"title"`
	Description string   `bson:"descprition" json:"desc"`
	Rating      int16    `bson:"rating" json:"rating"`
	Genre       []string `bson:"genre" json:"genre"`
}

func main() {

	var show Show
	coll := Connect()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// insertData(ctx, coll)

	err:= coll.FindOne(ctx, bson.M{"rating":4}).Decode(&show)
	if err != nil{
		log.Println("error fetching data", err)
	}

	jsonFormat,_ := json.Marshal(show)
	fmt.Println("got this shit",string(jsonFormat))
}
