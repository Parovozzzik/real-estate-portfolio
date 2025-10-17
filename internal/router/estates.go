package router

import (
    "net/http"

    "github.com/go-chi/chi/v5"

    "github.com/Parovozzzik/real-estate-portfolio/internal/database"
    "github.com/Parovozzzik/real-estate-portfolio/internal/handlers"
    "github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
)

func estatesRouter() http.Handler {
    db := database.GetDBInstance()
    estateRepository := repositories.NewEstateRepository(db)
    estateHandler := handlers.NewEstateHandler(estateRepository)

    r := chi.NewRouter()
    r.Get("/", estateHandler.GetEstates)
    r.Post("/", estateHandler.CreateEstate)

    return r
}
