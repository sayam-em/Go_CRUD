package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/sayam-em/Go_CRUD/Handlers"
)

func main()  {
	r := mux.NewRouter()

	r.HandleFunc("/user", handler.CreateUserHandler).Methods("POST")

	r.HandleFunc("/user", handler.GetUserHandler).Methods("GET")

	r.HandleFunc("/user", handler.PutUserHandler).Methods("PUT")

	r.HandleFunc("/user", handler.DeleteUserHandler).Methods("DELETE")


	log.Println("server is running on port :8090")
	log.Fatal(http.ListenAndServe(":8090", r))

	
}

