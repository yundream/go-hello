package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func h_hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello %s!", name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{name}", h_hello)
	log.Fatal(http.ListenAndServe(":8080", r))
}
