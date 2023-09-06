package main

/**************
   Functions are Generaly
   Placed here
****************/

func do_nothing() {
	/*do nothing */
}

// Check Error  \\
func checkError(err error) {
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
}

// Check Error as-Fatal \\
func checkError_asFatal(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
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

func get_data_from_db( /*specification of that data*/ /*) /*json {
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
