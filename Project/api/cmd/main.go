package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"				// CRUD Operation
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

func main()
{
	// chargement de variable global
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	
	/*
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
        panic(err)
	}*/

	// chargement de la base de données et de la collection ==> a changer en fonction du nom
	coll := client.Database("machine-L-DB").Collection("collect_L")
	title := "Machine Learning Data visualization based project Database"
	var result bson.M
	
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)
}
