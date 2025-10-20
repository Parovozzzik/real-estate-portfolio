package router

import (
    "net/http"

    "github.com/go-chi/chi/v5"

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

    r.Group(func(r chi.Router) {
        r.Use(JWTMiddleware)

        r.Get("/{id}/estates", userHandler.GetUserEstates)
    })

    return r
}
