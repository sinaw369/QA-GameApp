package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MySqlDB struct {
	db *sql.DB
}

func New() *MySqlDB {
	db, err := sql.Open("mysql", "gameapp:some_example@(localhost:3306)/gameapp_db")
	if err != nil {
		panic(fmt.Errorf("error creating database: %v", err))
	}
	//see import setting section
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySqlDB{db}
}
