package main

// ____ ENUMS

/**** Enum declaration
	type Season int

	const (
		Spring Season = iota + 1
		Summer
		Autumn
		Winter
	)
****/

/**** use Enum

//__
func (s Season) String() string {
	seasons := [...]string{"spring", "summer", "autumn", "winter"}
	if s < Spring || s > Winter {
		return fmt.Sprintf("Season(%d)", int(s))
	}
	return seasons[s-1]
}
//__

//__ Use to validate the value of a enum specially for user input data (if it match)
func (s Season) IsValid() bool {
	switch s {
	case Spring, Summer, Autumn, Winter:
		return true
	}
	return false
}
//__

//__ reading values
	mySeasons := []Season{Spring, Summer, Autumn, Winter, 6}

	for _, s := range mySeasons {
		fmt.Printf("season: %s, is valid: %t\n", s, s.IsValid())
	}
//__

****/

// - - - - - - - - - - - - - - -

// ____ API ENUMS

/**** Enum declaration
	type ApiError int

	const (
		Spring Season = iota + 1
		Summer
		Autumn
		Winter
	)
****/

// ____ CONSTANT
const (
	api_listen_port       int    = 2501
	api_base_url          string = ""
	api_serve_file_dir    string = "public"
	api_system_name       string = "machine-L" // can be Elyndo or anything else
	api_database_name_env string = "MACHINE_L_DB_NAME"
)

//const api_errors = slice

//______
