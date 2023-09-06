package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route // array of Route type

//_____

/**** 	Main Routes
_____
/api										| GET

/api/cmd/o-sign								| POST
/api/cmd/o-sign/conf						| POST

/api/cmd/o-connect							| POST

/api/cmd/query?user=;xxx=					| GET
/api/cmd/submit?user=; xxx=					| GET
_____

*****/

//_____

var api_routes = Routes{
	Route{
		"APIHomePage",
		"GET",
		"/api",
		apiHome_handler,
	},
	Route{
		"RegisterAPI",
		"POST",
		"/api/cmd/o-sign",
		regist_handler,
	},
	Route{
		"ConfirmRegistAPI",
		"POST",
		"/api/cmd/o-sign/conf",
		registConf_handler,
	},
	Route{
		"ConnectAPI",
		"POST",
		"/api/cmd/o-connect",
		connect_handler,
	},

	/*
		Route{
			"CityList",
			"GET",
			"/city",
			CityList,
		},

		Route{
			"CityDisplay",
			"GET",
			"/city/{id}",
			CityDisplay,
		},

		Route{
			"CityAdd",
			"POST",
			"/cityadd",
			CityAdd,
		},

		Route{
			"CityDelete",
			"GET",
			"/citydel/{id}",
			CityDelete,
		},
	*/
}
