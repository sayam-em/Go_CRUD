package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	Conn "github.com/sayam-em/Go_CRUD/DB"
	"github.com/sayam-em/Go_CRUD/Err"
	DbFunc "github.com/sayam-em/Go_CRUD/dbFunc"
)

type User struct {
    ID    int
    Name  string
    Email string
  }

  func GetUserHandler(w http.ResponseWriter, r *http.Request) {

    dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", Conn.DbUser, Conn.DbPass, Conn.DbName, Conn.Host, Conn.Port)

    db, err := sql.Open("postgres", dbInfo)
    if err != nil {
      Err.LogErr(err)
    }
  
    defer db.Close()

    // Get the 'id' parameter from the URL
    vars := mux.Vars(r)
    idStr := vars["id"]

    // Convert 'id' to an integer
    userID, err := strconv.Atoi(idStr)
    log.Println((userID))

    if err != nil {
      Err.LogErr(err)
    }


    // Call the GetUser function to fetch the user data from the database
    user, err := DbFunc.GetUser(db, userID)
    if err != nil {
      http.Error(w, "User not found", http.StatusNotFound)
      return
    }

    // Convert the user object to JSON and send it in the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
 }