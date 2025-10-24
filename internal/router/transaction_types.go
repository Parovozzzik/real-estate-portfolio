package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Parovozzzik/real-estate-portfolio/internal/database"
	"github.com/Parovozzzik/real-estate-portfolio/internal/handlers"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
)

func transactionTypesRouter() http.Handler {
	db := database.GetDBInstance()
	transactionTypeRepository := repositories.NewTransactionTypeRepository(db)
	transactionTypeHandler := handlers.NewTransactionTypeHandler(transactionTypeRepository)

	r := chi.NewRouter()
	r.Get("/", transactionTypeHandler.GetTransactionTypes)
	r.Post("/", transactionTypeHandler.CreateTransactionType)
	r.Put("/{id}", transactionTypeHandler.UpdateTransactionType)

	return r
}
