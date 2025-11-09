package main

import (
	"fmt"
	"net/http"

	"github.com/Parovozzzik/real-estate-portfolio/internal/database"
	"github.com/Parovozzzik/real-estate-portfolio/internal/router"
	"github.com/Parovozzzik/real-estate-portfolio/pkg/logging"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	defer database.CloseDB()

	logger.Println("router initializing")
	r := router.GetRouter()

	logger.Println("start application")
	err := http.ListenAndServe(":8085", r)
	if err != nil {
		fmt.Println("Server failed to start: %v", err)
	}
}
