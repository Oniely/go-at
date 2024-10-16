package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe("localhost:8081", myRouter))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	fmt.Println("GO ORM")

	InitialMigration()

	handleRequests()
}
