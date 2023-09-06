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
	
	"github.com/gorilla/mux"
	
	//...
	
)

// ____ Enums


// ____ Structs


/*
api_error struct {
	
}
*/

// ____ Constants
const {
	api_launch_port = 2501
}
	

//const api_errors = slice

// ______

func main() {

	//getdata()

	//  _____ step 1: loading template

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveTemplate)


	// _____ step 2: listen http socket
	
	
	fmt.Println("Starting the server on :"+ string(api_launch_port))

	err := http.ListenAndServe(":"+string(api_launch_port), nil)
	if err != nil {
		log.Fatal(err)
	}
	else {
		fmt.Println("Server Start on port 2501, Successfully!")
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
		log.Fatal("Sorry! You must set your 'MONGODB_URI' environmental variable."+
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

	/*	
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
        panic(err)
	}
	*/

	//_____ step 2: Recovering Database Name

	db := os.Getenv("MACHINE_L_DB_NAME")
	if db == "" {
		log.Fatal("Sorry! You must set your Database's Name in a called 'MACHINE_L_DB_NAME' environmental variable."+
		 					"\nSee \t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")						
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


/// ______ FUNCTIONS ________
/*
func registration_process() {
	//
}

func registration_confirm_process() {
	//
}

func connection_process() {
	//
}

func query_data_process() {
	//
}

func submit_data_process() {
	//
}


func modify_data_process() {
	//
}


func delete_data_process() {
	//
}

//____

func get_page_header_content_type() {
	//
}

func get_received_API_URL(){
	//
}

func get_received_json(){
	//
}

func treatPurpose__registration(){
	//
}

func isData_match_required_format(json string, process_enum string ) bool { // process is a enum not a string
	//
}

func save_data_to_db(json string) {
	//
}

func send_client_response_wrong(process_enum string) json {
	//
}

func send_client_response_success(process_enum string) json {
	//
}

func get_data_from_db( /*specification of that data*/) /*json {
	//
}

func get_api_token_from_db() string {
	//
}

func generate_new_token() string {
	//
}

func isToken_used() bool {
	//
}

//____ security

func hash_pass(pass string) string {
	//
}


func retrieve_hashed_pass(pass string) string {
	//
}

//____ from_ Functions
/*
func from_url_to_process_enum() process_enum {
	//
}

func from_json_to_slice(json string) slice {
	//
}

func from_slice_to_json(slice string) json {
	//
}


func from_json_data_inside_to_prop() {
	//
}

func from_json_data_inside_to_string() {
	//
}

func from_json_data_inside_to_int()
{
	//
}
*/