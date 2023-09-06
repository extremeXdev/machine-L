package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/go-playground/validator/v10"
)

//______

// __
func apiHome_handler(w http.ResponseWriter, r *http.Request) {
	/*
		fmt.Fprintln(w, "Welcome to the "+string(api_system_name)+" API!\n"+
			"--> add '/cmd/' to access online command...\n"+
			"For more infos see our documentation here at 'api/doc/'")
	*/

	//  _____ serving template for http system

	fs := http.FileServer(http.Dir("./" + string(api_serve_file_dir))) // public

	http.Handle("/"+string(api_serve_file_dir)+"/",
		http.StripPrefix("/"+api_serve_file_dir+"/", fs)) //'public' OR 'static'

	pattern := "/"

	http.HandleFunc(pattern, serveTemplate)

}

// Serve Template \\
func serveTemplate(w http.ResponseWriter, r *http.Request) { // rw ;  req

	//regex := regexp.MustCompile("favicon.ico")

	fmt.Println(r.URL.Path)
	fmt.Println("\n =========>\n")
	lp := filepath.Join("templates", "layout.html")

	var fp string

	if r.URL.Path == "/" {
		fp = filepath.Join("templates", filepath.Clean("/welcome.html")) // let set as default welcome page
	} else {
		fp = filepath.Join("templates", filepath.Clean(r.URL.Path)) // the uri is right oriented to do so
	}

	tmpl, _ := template.ParseFiles(lp, fp)

	tmpl.ExecuteTemplate(w, "layout", nil)
}

//\\\\

// registration _handler \\
func regist_handler(w http.ResponseWriter, r *http.Request) {
	var registAcc APIRegistData

	//:::: Read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	//::::

	//:::: Convert the JSON in the request to a Go type
	if err := json.Unmarshal(body, &registAcc); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422) // can't process!

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	//::::

	//:::: Write to the database
	addResult := dbUserAdd(user)
	//::::

	//:::: Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(addResult)
	//::::
}

// registration confirmation _handler \\
func registConf_handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userName := vars["api_user_name"]
}

// connection _handler \\
func connect_handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userName := vars["api_user_name"]
}

//______

//
//
//

// __
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // transform received data in a map called vars
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

//__

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productID := vars["id"]

	log.Printf("Product ID:%v\n", productID)

	fileName := productID + ".html"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Printf("no such product")

		fileName = "invalid.html"
	}

	http.ServeFile(w, r, fileName)
}

//___________________________

//___________________________

func UserList(w http.ResponseWriter, r *http.Request) {
	// Query the database
	jsonUsers := dbUserList()

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUsers)
}

func UserDisplay(w http.ResponseWriter, r *http.Request) {
	// Get URL parameter with the user ID to search for
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	// Query the database
	jsonUser := dbUserDisplay(userId)

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)
}

func UserAdd(w http.ResponseWriter, r *http.Request) {
	var user User

	// Read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Convert the JSON in the request to a Go type
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422) // can't process!
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// Write to the database
	addResult := dbUserAdd(user)

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(addResult)
}

func UserDelete(w http.ResponseWriter, r *http.Request) {

	// Get URL parameter with the user ID to delete
	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["id"], 10, 64)

	// Query the database
	deleteResult := dbUserDelete(userId)

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(deleteResult)
}
*/