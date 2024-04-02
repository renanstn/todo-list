package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *sql.DB
var gormDB *gorm.DB

func ConnectDatabase() {
	database_url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", database_url+"?sslmode=disable")
	if err != nil {
		fmt.Println("There is an error while connecting to the database:", err)
		panic(err)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
	}
	// Initialize ORM
	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: Db,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("There is an error while connecting to ORM:", err)
		panic(err)
	} else {
		fmt.Println("ORM successfully connected!")
	}
}

func SetupDatabase() {
	gormDB.AutoMigrate(&Todo{})
}
