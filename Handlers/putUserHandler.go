package handler

import (
	"database/sql"
	"errors"

	Conn "github.com/sayam-em/Go_CRUD/DB"
)
func putUserHandler() {

	db, err := sql.Open(Conn.DbDriver)
	errors.Is()

}
