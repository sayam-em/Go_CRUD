package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/sayam-em/Go_CRUD/Err"
	"github.com/sayam-em/Go_CRUD/DB"
	"github.com/sayam-em/Go_CRUD/DB"
)

func createUserHandler() {

	db, err := sql.Open(Conn.dbDriver,Conn.dbUser+":"+Conn.dbPass+"@/"+Conn.dbName )
	if err != nil {
		Err.LogErr(err)
	}

	defer db.Close()

	var user User
	json.NewDecoder(r.Body, user.Email)

	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "user created")




}
