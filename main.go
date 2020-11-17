package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	fmt.Println(reflect.TypeOf(r))
	log.Fatal(http.ListenAndServe(":8000", r))
}
