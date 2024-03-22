package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sayam-em/Go_CRUD/DB"
	"github.com/sayam-em/Go_CRUD/Err"
	DbFunc "github.com/sayam-em/Go_CRUD/dbFunc"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", Conn.DbUser, Conn.DbPass, Conn.DbName, Conn.Host, Conn.Port)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		Err.LogErr(err)
	}

	defer db.Close()

	vars := mux.Vars(r)

	idStr := vars["id"]

	userId, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	user := DbFunc.DeleteUser(db, userId)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User deleted successfully")
	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}
