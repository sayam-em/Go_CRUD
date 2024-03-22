package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	Conn "github.com/sayam-em/Go_CRUD/DB"
	"github.com/sayam-em/Go_CRUD/Err"
	DbFunc "github.com/sayam-em/Go_CRUD/dbFunc"
)

func PutUserHandler(w http.ResponseWriter, r *http.Request) {

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", Conn.DbUser, Conn.DbPass, Conn.DbName, Conn.Host, Conn.Port)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		Err.LogErr(err)
	}

	defer db.Close()

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	DbFunc.UpdateUser(db, user.ID, user.Name, user.Email)

	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User updated successfully")

}
