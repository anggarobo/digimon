package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Database struct {
	User     string
	Host     string
	Port     string
	Password string
	Name     string
}

var SQL *sql.DB

func (d *Database) Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	d.User = os.Getenv("DB_USER")
	d.Host = os.Getenv("DB_HOST")
	d.Port = os.Getenv("DB_PORT")
	d.Password = os.Getenv("DB_PASSWORD")
	d.Name = os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetConnMaxLifetime(time.Minute * 60)
	fmt.Println("Connected to Mysql database!")

	SQL = db
}
