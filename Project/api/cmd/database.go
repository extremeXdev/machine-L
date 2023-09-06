package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"strconv"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	//
	//...
	//...
	//___ sql database need
	/*
		"database/sql"
		"encoding/json"
		"fmt"
		"log"

		_ "github.com/go-sql-driver/mysql"
	*///
)

var database_cli *mongo.Client
var database_noSQL *mongo.Database

// Connect to the "world" database
func dbConnect_noSQL() {

	// _____ step 0: Recovering Mongo DB URI

	// loading global var
	if err := godotenv.Load(); err != nil {
		log.Println("No .env (environment) file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Sorry! You must set your 'MONGODB_URI' environmental variable." +
			"\nSee \t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	// _____ step 1: Testing connection

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	database_cli = client

	//_____ step 2: Recovering Database Name

	db_name := os.Getenv(string(api_database_name_env))
	if db_name == "" {
		log.Fatal("Sorry! You must set your Database's Name in a called 'MACHINE_L_DB_NAME' environmental variable." +
			"\nSee \t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		return
	}

	// _____ step 3: establishing database connection

	database_noSQL = database_cli.Database(db_name)
	log.Println("Connected Successfully to NOSQL Database")
}

//____________________________

// ping database no-SQL \\
func pingDB_noSQL() bool {
	rsl := true

	if err := database_cli.Ping(context.TODO(), readpref.Primary()); err != nil {
		rsl = false
		//panic(err)
	}

	return rsl
}

//____________________________

var database *sql.DB

// Connect to the "world" database
func dbConnect() {
	// replace "root" and "password" with your database login credentials
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/world")
	if err != nil {
		log.Println("Could not connect!")
	}
	database = db
	log.Println("Connected.")
}

// Find all User and return as JSON
func dbUserList() []byte {
	//...
}

// Find a single User based on ID and return as JSON
func dbUserDisplay(id int) []byte {
	//...
}

// Create a new User based on the information supplied
func dbUserAdd(User User) []byte {
	//...
}

// Delete the User with the supplied ID
func dbUserDelete(id int64) []byte {
	//...
}
