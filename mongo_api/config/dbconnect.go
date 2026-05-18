package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var ShowClcn *mongo.Collection

func Connect(){
	err := godotenv.Load(".env");


	if err != nil{
		log.Fatalln("error loading the env file", err)
	}

	cntn_url := os.Getenv("MONGO_URI")
	db_name := os.Getenv("DATABASE_NAME")

	cntn_opts := options.Client().ApplyURI(cntn_url)

	mongo_client, err := mongo.Connect(context.Background(), cntn_opts)

	if err != nil{
		log.Fatalln("error creating mongo client ", err)
	}else{
		fmt.Println("connected to the mongo db at :", cntn_url)
	}


	err = mongo_client.Ping(context.Background(), nil)
	if err != nil{
		log.Fatalln("unable to ping database", err)
	}

	DB = mongo_client.Database(db_name)

	ShowClcn = DB.Collection("shows")

}