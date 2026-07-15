/*
 * Now let's do something that we actually use in the development
 * of division online... Expanding a very simple API.
 *
 * The code below creates an API server that exposes a single
 * "HTTP GET" API endpoint (i.e. "/hi/{name}") that returns a
 * greeting message on access.
 *
 * Example:
 *  If you run the program as is, and then go to
 *  localhost:1234/hi/alex in your browser, it should display the
 *  message "Hi, alex!". You can replace the name alex with any
 *  name and it will greet you accordingly...
 *
 * To do:
 *  Try to understand the code.
 *  Implement a goodbye "GET" endpoint (i.e. "/bye/{name}") that
 *  says "Goodbye, {name}!".
 *  Try implementing a POST endpoint (or DELETE, or PUT...), for
 *  example a "POST /echo" endpoint that reads the request body
 *  and sends it right back in the response.
 *
 * This is basically 90% of what you are expected to do with Go
 * for the purpose of this project... API's (and some Websockets).
 * */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func getHiGreeting(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	fmt.Fprintf(w, "Hi, %s!", name)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hi/{name}", getHiGreeting)
	// mux.HandleFunc("GET /bye/{name}", getByeGreeting) // implement getByeGreeting
	// mux.HandleFunc("POST /echo", postEcho) // implement postEcho

	log.Println("listening on 127.0.0.1:1234")
	if err := http.ListenAndServe("127.0.0.1:1234", mux); err != nil {
		log.Fatal(err)
	}
}
