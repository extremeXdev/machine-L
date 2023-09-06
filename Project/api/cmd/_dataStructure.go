package main

// ____ STRUCTS

/**** struct declaration
	//__
	type Fruit struct {
		name string
	}
	//__

	//__
	type User struct {
	    name string
	}

	type Service struct {
	        name string
	    user User
	}
	//__

	//__
	type Person struct {
	    name string
	}

	type StoreKeeper struct {
	    Person
	}
 	//__

	//__ Declaring and initializing an Array of Structs
		//var arrayOfStruct []StructName     // --> We simply declare an array with the type of that struct.
		var students []Student
	//__

*****/

/**** Struct use
	//__
		var apple = Fruit{"Apple"} // struct literal syntax
		fmt.Println(apple)      // prints {Apple}
	//__

	//__
		var banana = new(Fruit)
	    banana.name = "Banana"
	    fmt.Println(banana)      // prints &{Banana}
    //__

    //__
	    var mango = &Fruit{"Mango"}
	    fmt.Println(mango)      // prints &{Mango}
    //__

    //__
    	google := Service{
    	    name: "Google",
    	    user: User{
    	        name: "John Doe",
    	    },
    	}

    	// accessing from nested struct
    	fmt.Println(google.user.name)  // prints "John Doe"
    //__

    //__
	    var p = StoreKeeper{}
	    p.Person = Person{"Jane Doe"}

	    // access the field as normal field
	    fmt.Println(p.name)                   // prints "Jane Doe"
    //__

    //__
	    // declare array
	    var students []Student

	    s := Student{"Kate"}
	    kate := append(students, s)
	    fmt.Println(kate)             // [{Kate}]
    //__
/****/

//_____

type City struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country"`
	District    string `json:"district"`
	Population  int    `json:"pop"`
}

type Cities []City

//_______________________________________

// ____ API STRUCT

type APIOwner struct {
	name  string
	state string
}

type APIUser struct {
	name  string
	pass  string
	mtk   string // main token
	email string

	owner APIOwner

	tokens []string
}

// ____ Registration

type APIRegistData struct {
	api_userId string // user ID

	api_username    string // user name
	api_pass        string // pass
	api_dev_email   string // dev email
	api_owner_name  string // owner name
	api_owner_state string // owner state
}

type APIRegistData_RSP struct {
	status   int    // status
	rsp_code string // response code
	rsp_desc string // response description
}

//____ Registration Confirmation

type APIRegistConfData struct {
	api_username string // 1
	api_pass     string // 2
}

type APIRegistConfData_RSP struct {
	token    string // the main token as a response
	rsp_code int
	rsp_desc string
}

//____ Connection

type APIConnData struct {
	api_username string // 1
	api_pass     string // 2
}

type APIConnData_RSP struct {
	token    string // the main token as a response
	rsp_code int
	rsp_desc string
}

//______

// ____ Query Data

//
//
//

//______
