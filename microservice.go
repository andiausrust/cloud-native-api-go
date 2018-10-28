package main

import (
	"cloud-native-api-go/api"
	"fmt"
	"net/http"
	"os"
)

func main () {
	http.HandleFunc("/", index)   		// "index" = handler function listening to the specified path
	http.HandleFunc("/api/echo", echo)

	http.HandleFunc("/api/books", api.BooksHandleFunc)

	http.ListenAndServe(port(), nil)
}

// it is good practice to specify a port function to make port configurable
// port variable has to be set: "export PORT=8080" in Terminal
func port() string{
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello world cloud native Go")
}

func echo(w http.ResponseWriter, r *http.Request){
	message := r.URL.Query()["message"][0]  	// extract the message query parameter from our URL (first parameter [0])
	fmt.Fprint(w, message)
}
