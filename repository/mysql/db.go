package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Config struct {
	UserName string
	Password string
	Port     int
	Host     string
	Scheme   string // db name

}
type MySqlDB struct {
	db     *sql.DB
	Config Config
}

func New(config Config) *MySqlDB {
	//parseTime  ture changes the output type if DATE and DATETIME values to time.time
	// instead of []byte / string
	// The date or datetime like 0000-00-00 00:00:00 is converted into zero value of time.time
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		config.UserName, config.Password, config.Host, config.Port, config.Scheme))
	if err != nil {
		panic(fmt.Errorf("error creating database: %v", err))
	}
	//see import setting section
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySqlDB{db: db, Config: config}
}
