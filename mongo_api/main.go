package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func Connect() *mongo.Collection{
	err:=godotenv.Load(".env")

	if err != nil{
		log.Println("error occured while loadong env", err)
	}

	connection_url := os.Getenv("MONGO_URI")

	conn_options := options.Client().ApplyURI(connection_url)

	mongo_client , err := mongo.Connect(context.Background(), conn_options)

	if err != nil{
		log.Println("error occured while creating mongo client ", err)
	}


	err = mongo_client.Ping(context.Background(), nil)
	if err != nil{
		log.Println("unable to ping the ")
	}else{
		fmt.Println("connected to mongodb")
	}

	collection := mongo_client.Database("e-commerse").Collection("test_collection")

	return  collection
}


func main(){
	Connect()
}