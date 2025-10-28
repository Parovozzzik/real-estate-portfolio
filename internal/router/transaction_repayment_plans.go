package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Parovozzzik/real-estate-portfolio/internal/database"
	"github.com/Parovozzzik/real-estate-portfolio/internal/handlers"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
)

func transactionRepaymentPlansRouter() http.Handler {
	db := database.GetDBInstance()
	transactionRepaymentPlanRepository := repositories.NewTransactionRepaymentPlanRepository(db)
	transactionRepaymentPlanHandler := handlers.NewTransactionRepaymentPlanHandler(transactionRepaymentPlanRepository)

	r := chi.NewRouter()
	r.Get("/", transactionRepaymentPlanHandler.GetTransactionRepaymentPlans)
	r.Post("/", transactionRepaymentPlanHandler.CreateTransactionRepaymentPlan)
	r.Put("/{id}", transactionRepaymentPlanHandler.UpdateTransactionRepaymentPlan)

	return r
}
