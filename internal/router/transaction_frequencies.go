package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Parovozzzik/real-estate-portfolio/internal/database"
	"github.com/Parovozzzik/real-estate-portfolio/internal/handlers"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
)

func transactionFrequenciesRouter() http.Handler {
	db := database.GetDBInstance()
	transactionFrequencyRepository := repositories.NewTransactionFrequencyRepository(db)
	transactionFrequencyHandler := handlers.NewTransactionFrequencyHandler(transactionFrequencyRepository)

	r := chi.NewRouter()
	r.Get("/", transactionFrequencyHandler.GetTransactionFrequencies)
	r.Post("/", transactionFrequencyHandler.CreateTransactionFrequency)
	r.Put("/{id}", transactionFrequencyHandler.UpdateTransactionFrequency)

	return r
}
