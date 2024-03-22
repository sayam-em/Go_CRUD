package main

import (
	"log"
	"net/http"
    "github.com/gorilla/mux"

)

func main()  {
	r := mux.NewRouter()

	r.HandleFunc("/user", createUserHandler).Methods("POST")

	r.HandleFunc("/user", getUserHandler).Methods("GET")

	r.HandleFunc("/user", putUserHandler).Methods("PUT")

	r.HandleFunc("/user", deleteUserHandler).Methods("DELETE")


	log.Println("server is running on port :8090")
	log.Fatal(http.ListenAndServe(":8090", r))

	
}

