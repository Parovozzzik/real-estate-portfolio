package router

import (
	"github.com/Parovozzzik/real-estate-portfolio/internal/services"
	"net/http"

	chi "github.com/go-chi/chi/v5"

	"github.com/Parovozzzik/real-estate-portfolio/internal/database"
	"github.com/Parovozzzik/real-estate-portfolio/internal/handlers"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
)

func usersRouter() http.Handler {
	db := database.GetDBInstance()
	userRepository := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepository)

	estateValueRepository := repositories.NewEstateValueRepository(db)
	transactionRepository := repositories.NewTransactionRepository(db)

	estateValuesService := services.NewEstateValuesService(estateValueRepository, transactionRepository)

	r := chi.NewRouter()
	r.Get("/", userHandler.GetUsers)
	r.Post("/login", userHandler.LoginUser)
	r.Post("/registration", userHandler.RegistrationUser)
	r.Post("/refresh-token", userHandler.RefreshToken)

	r.Group(func(r chi.Router) {
		r.Use(JWTMiddleware)
		r.Use(UserIdMatchMiddleware)

		r.Get("/{user-id}", userHandler.GetUserById)
		r.Put("/{user-id}/profile", userHandler.UpdateUser)
		r.Get("/{user-id}/estates", userHandler.GetUserEstates)
		r.Get("/{user-id}/transactions", userHandler.GetUserTransactions)
		r.Get("/{user-id}/estates/{estate-id}", userHandler.GetUserEstate)
		r.Post("/{user-id}/estates/{estate-id}/values/filter", userHandler.GetUserEstateValues)
		r.Post("/{user-id}/estates/{estate-id}/values/recalculate", estateValuesService.Recalculate)
		r.Post("/{user-id}/estates", userHandler.CreateEstate)
		r.Put("/{user-id}/estates/{estate-id}", userHandler.UpdateEstate)
		r.Post("/{user-id}/estates/{estate-id}/transactions/filter", userHandler.GetUserEstateTransactions)
	})

	return r
}
