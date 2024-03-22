package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	Conn "github.com/sayam-em/Go_CRUD/DB"
	Err "github.com/sayam-em/Go_CRUD/Err"
	DbFunc "github.com/sayam-em/Go_CRUD/dbFunc"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", Conn.DbUser, Conn.DbPass, Conn.DbName, Conn.Host, Conn.Port)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		Err.LogErr(err)
		http.Error(w, "failed to connect to database", http.StatusInternalServerError)
		return
	}

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	err = DbFunc.CreateUser(db, user.Name, user.Email)
	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "user created")
}
