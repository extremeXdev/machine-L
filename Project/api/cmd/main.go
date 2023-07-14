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
)

func main() {

	//getdata()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveTemplate)
	fmt.Println("Starting the server on :2501")
	err := http.ListenAndServe(":2501", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	//regex := regexp.MustCompile("favicon.ico")
	fmt.Println(r.URL.Path)
	fmt.Println("\n =========>\n")
	lp := filepath.Join("templates", "layout.html")
	var fp string
	if r.URL.Path == "/" {
		fp = filepath.Join("templates", filepath.Clean("/welcome.html"))
	} else {
		fp = filepath.Join("templates", filepath.Clean(r.URL.Path))
	}
	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func getdata() {

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
	// chargement de la base de données et de la collection ==> a changer en fonction du nom
	coll := client.Database("Test").Collection("Nom de la collection")
	title := "Example de titre"
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
