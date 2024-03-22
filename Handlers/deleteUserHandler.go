package handler

import (
	"database/sql"
	"errors"

	Conn "github.com/sayam-em/Go_CRUD/DB"
)
func deleteUserHandler() {

	db, err := sql.Open(Conn.DbDriver)
	errors.Is()

}
