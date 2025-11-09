package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Parovozzzik/real-estate-portfolio/pkg/config"
)

var (
	dbInstance *sql.DB
	once       sync.Once
)

func GetDBInstance() *sql.DB {
	once.Do(func() {
		cfg := config.GetConfig()

		dsn := cfg.MySql.Username + ":" + cfg.MySql.Password + "@tcp(" + cfg.MySql.Host + ":" + cfg.MySql.Port + ")/" + cfg.MySql.Database

		var err error
		dbInstance, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Error opening database connection: %v", err)
		}

		err = dbInstance.Ping()
		if err != nil {
			log.Fatalf("Error connecting to the database: %v", err)
		}

		fmt.Println("Database connection established successfully!")
	})
	return dbInstance
}

func CloseDB() {
	if dbInstance != nil {
		err := dbInstance.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			fmt.Println("Database connection closed.")
		}
	}
}
