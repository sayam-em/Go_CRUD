// Conn package
package Conn

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

var (
	DbUser = "postgres"
	DbPass = "password"
	DbName = "postgres"
	Host   = "localhost"
	Port   = "5432"
)

func init() {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", DbUser, DbPass, DbName, Host, Port)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
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
