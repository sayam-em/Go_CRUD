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
	DbUser   = "dbUser"
	DbPass   = "dbPass"
	DbName   = "gocrud_app"
	Host     = "localhost"
	Port     = "5432"
)

func init() {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s dbdriver=%s host=%s port=%s sslmode=disable", DbUser, DbPass, DbName, DbDriver, Host, Port)

	db, err := NewDb(dbInfo)
	Err.LogErr(err)

	rows, err := db.Query("SELECT * from users")
	Err.LogErr(err)

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
