package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	Conn "github.com/sayam-em/Go_CRUD/Conn"
	Err "github.com/sayam-em/Go_CRUD/Err"
)

func createUserHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open(Conn.DB.DbDriver, Conn.dbUser+":"+Conn.dbPass+"@/"+Conn.dbName)
	if err != nil {
		Err.LogErr(err)
	}

	defer db.Close()

	var user User
	json.NewDecoder(r.Body).Decode(user.Email)

	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "user created")

}
