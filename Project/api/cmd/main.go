package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"strconv"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/gorilla/mux"
	//...
)

// ______

func main() {

	/*
		//router.Schemes("https")

		r.Headers("Content-Type", "application/json")
		//r.Headers("X-Requested-With", "XMLHttpRequest")

		// setting API exchange method
		r.Methods("GET", "POST")

		//	API Home
		//	/
		r.HandleFunc("/", apiHome_handler)

		//	Registering URI suffix match
		//	/api/cmd/o-sign
		r.HandleFunc("/api/cmd/o-sign", regist_handler)

		//	Confirmation URI suffix match
		//	/api/cmd/o-sign/conf
		r.HandleFunc("/api/cmd/o-sign/conf", registConf_handler)
	*/
	//
	//
	//r := mux.NewRouter() // mux http router

	/*
	    r.HandleFunc("/", HomeHandler)
	    r.HandleFunc("/products", ProductsHandler)
	    r.HandleFunc("/articles", ArticlesHandler)


	   	r.HandleFunc("/products/{key}", ProductHandler)
	   	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	   	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	*/
	/*
		// Matches a dynamic subdomain.
		r.Host("{subdomain:[a-z]+}.example.com")
	*/

	//http.Handle("/", r)

	// -------   -------   -------

	//  _____ step 1: API launch Socket

	//:::: Socket by default Setting
	api_socket := api_launch_port
	//::::

	//:::: Socket by Command line -- user can optionaly send port
	if len(os.Args) >= 2 {
		port := int(os.Args[1]) // -- argument 1 must be the port

		if port > 100 {
			api_socket = port
			fmt.println(" API going to be launched on port: " +
				string(api_socket))
		} else {
			fmt.println(" The port given through commande-line\n " +
				"isn't accepted! -- default set to : " + string(api_socket))
		}
	}
	//::::

	//:::: Establishing The Rooter
	router := NewAPIRouter()
	//::::

	// connect to database always before Listen And Serve

	//:::: Connecting MongoDB
	dbConnect_noSQL()
	//::::

	//:::: ???
	getdata()
	//::::

	//
	// _____ step 3: listen http socket

	fmt.Println("Starting the server on :" + string(api_socket) + " ...")

	err := http.ListenAndServe(":"+string(api_socket), router)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Server Started on port %d, Successfully!", api_socket)
	}

}

/// - - - - - - - - - - - - - - -  Functions  - - - - - - - - - - - - - - - \\\

//_________

/* call in main file or above
check := func(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
*/

// // Get Data Process  \\\\
func getdata(db *mongo.Database) {

	// _____ step 1: loading document collection 'x' data as willing

	coll := db.Collection("collect_L")

	title := "Example de titre"

	var result bson.M

	err = coll.FindOne(context.TODO(), bson.D{{"_title_", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}

	// _____ step 2: json data loading start here

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData) // display the json content

}
