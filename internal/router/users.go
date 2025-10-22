package router

import (
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

	r := chi.NewRouter()
	r.Get("/", userHandler.GetUsers)
	r.Post("/login", userHandler.LoginUser)
	r.Post("/registration", userHandler.RegistrationUser)
	r.Post("/refresh-token", userHandler.RefreshToken)

	r.Group(func(r chi.Router) {
		r.Use(JWTMiddleware)

		r.Get("/{user-id}", userHandler.GetUserById)
		r.Put("/{user-id}/profile", userHandler.UpdateUser)
		r.Get("/{user-id}/estates", userHandler.GetUserEstates)
		r.Get("/{user-id}/estates/{estate-id}", userHandler.GetUserEstate)
	})

	return r
}
