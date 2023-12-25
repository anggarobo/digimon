package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	User     string
	Host     string
	Port     string
	Password string
	Name     string
}

var DB *gorm.DB

func (d Database) Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	d.User = os.Getenv("DB_USER")
	d.Host = os.Getenv("DB_HOST")
	d.Port = os.Getenv("DB_PORT")
	d.Password = os.Getenv("DB_PASSWORD")
	d.Name = os.Getenv("DB_NAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
	)

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to Mysql database!")

	DB = db
}
