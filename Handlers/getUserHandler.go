package handler

import (
	"database/sql"
	"errors"

	Conn "github.com/sayam-em/Go_CRUD/DB"
	Err "github.com/sayam-em/Go_CRUD/Err"

)


type User Struct {

	ID int
	Name string
	Email string

}


func getUserHandler() {

	db, err := sql.Open(Conn.dbDriver,Conn.dbUser+":",conn.dbPass+"@/",conn.dbName)

	if err = != nil {
		Err.LogErr(err)
		
		
	}

	defer db



}
