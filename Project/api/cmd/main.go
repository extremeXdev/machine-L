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

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	
	//...
	
)

func main() {

	//getdata()

	//  _____ step 1: loading template

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveTemplate)


	// _____ step 2: listen http socket
	
	fmt.Println("Starting the server on :2501")

	err := http.ListenAndServe(":2501", nil)
	if err != nil {
		log.Fatal(err)
	}
	
}


//// Functions

func serveTemplate(w http.ResponseWriter, r *http.Request) {

	//regex := regexp.MustCompile("favicon.ico")

	fmt.Println(r.URL.Path)
	fmt.Println("\n =========>\n")
	lp := filepath.Join("templates", "layout.html")

	var fp string

	if r.URL.Path == "/" {
		fp = filepath.Join("templates", filepath.Clean("/welcome.html"))   // let set as default welcome page
	} else {
		fp = filepath.Join("templates", filepath.Clean(r.URL.Path))		   // the uri is right oriented to do so
	}

	tmpl, _ := template.ParseFiles(lp, fp)
	
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func getdata() {

	// _____ step 0: Recovering Mongo DB URI

	// loading global var
	if err := godotenv.Load(); err != nil {
		log.Println("No .env (environment) file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Sorry! You must set your 'MONGODB_URI' environmental variable.\
			\nSee \t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
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

	//_____ step 2: Recovering Database Name

	db := os.Getenv("MACHINE_L_DB_NAME")
	if db == "" {
		log.Fatal("Sorry! You must set your Database's Name in a called 'MACHINE_L_DB_NAME' environmental variable.\
		 					\nSee \t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")						
		return
	}

	// _____ step 3: loading collection 'x' data as willing

	coll := client.Database(db).Collection("collect_L")
	
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

	// _____ step 4: json data loading start here
	
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)  // display the json content

}
