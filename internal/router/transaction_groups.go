package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Parovozzzik/real-estate-portfolio/internal/database"
	"github.com/Parovozzzik/real-estate-portfolio/internal/handlers"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
)

func transactionGroupsRouter() http.Handler {
	db := database.GetDBInstance()
	transactionGroupRepository := repositories.NewTransactionGroupRepository(db)
	transactionGroupHandler := handlers.NewTransactionGroupHandler(transactionGroupRepository)

	r := chi.NewRouter()
	r.Get("/{estate-id}", transactionGroupHandler.GetTransactionGroups)
	r.Post("/", transactionGroupHandler.CreateTransactionGroup)
	r.Put("/{id}", transactionGroupHandler.UpdateTransactionGroup)

	return r
}
