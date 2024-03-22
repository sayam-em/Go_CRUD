package Conn

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sayam-em/Go_CRUD/Err"
)

type DB struct {
	*sql.DB
}

var (
	DbDriver = "postgres"
	dbUser   = "dbUser"
	dbPass   = "dbPass"
	dbName   = "gocrud_app"
	host     = "localhost"
	port     = "5432"
)

func init() {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s dbdriver=%s host=%s port=%s sslmode=disable", dbUser, dbPass, dbName, DbDriver, host, port)

	db, err := NewDb(dbInfo)
	Err.CheckErr(err)

	rows, err := db.Query("SELECT * from users")
	Err.CheckErr(err)

	fmt.Println(rows)
}

func NewDb(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
