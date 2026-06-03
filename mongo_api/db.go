package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)




func DBconnect() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil{
		log.Fatalln("failed to load env", err)
	}


	con_url := os.Getenv("MONGO_URI")

	if con_url == ""{
		log.Fatalln("the url is not loaded please fix the .env")
	}


	client , err := mongo.Connect(options.Client().ApplyURI(con_url))

	if err != nil{
		log.Fatalln("failed to connect to clinet", err)
	}


	ctxTime , cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()
	if err = client.Ping(ctxTime, nil); err != nil{
		log.Fatalln("failed to ping the mongo")
	}

	return client

}