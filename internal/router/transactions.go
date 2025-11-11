package router

import (
	"github.com/Parovozzzik/real-estate-portfolio/internal/services"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Parovozzzik/real-estate-portfolio/internal/database"
	"github.com/Parovozzzik/real-estate-portfolio/internal/handlers"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
)

func transactionsRouter() http.Handler {
	db := database.GetDBInstance()
	transactionRepository := repositories.NewTransactionRepository(db)
	transactionHandler := handlers.NewTransactionHandler(transactionRepository)

	transactionGroupRepository := repositories.NewTransactionGroupRepository(db)
	transactionTypeRepository := repositories.NewTransactionTypeRepository(db)
	transactionFrequencyRepository := repositories.NewTransactionFrequencyRepository(db)
	transactionRepaymentPlanRepository := repositories.NewTransactionRepaymentPlanRepository(db)
	transactionGroupSettingRepository := repositories.NewTransactionGroupSettingRepository(db)

	transactionService := services.NewTransactionService(
		transactionRepository,
		transactionGroupRepository,
		transactionTypeRepository,
		transactionFrequencyRepository,
		transactionRepaymentPlanRepository,
		transactionGroupSettingRepository,
	)

	r := chi.NewRouter()
	r.Get("/", transactionHandler.GetTransactions)
	r.Get("/{estate-id}/{group-id}/{type-id}", transactionHandler.GetTransactions)
	r.Post("/", transactionService.CreateTransaction)
	r.Put("/", transactionService.UpdateTransaction)
	r.Put("/{transaction-id}", transactionHandler.UpdateTransaction)
	r.Delete("/{transaction-id}", transactionHandler.DeleteTransaction)

	return r
}
