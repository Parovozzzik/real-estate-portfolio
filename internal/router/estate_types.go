package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Parovozzzik/real-estate-portfolio/internal/database"
	"github.com/Parovozzzik/real-estate-portfolio/internal/handlers"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
)

func estateTypesRouter() http.Handler {
	db := database.GetDBInstance()
	estateTypeRepository := repositories.NewEstateTypeRepository(db)
	estateTypeHandler := handlers.NewEstateTypeHandler(estateTypeRepository)

	r := chi.NewRouter()
	r.Get("/", estateTypeHandler.GetEstateTypes)
	r.Post("/", estateTypeHandler.CreateEstateType)
	r.Put("/{id}", estateTypeHandler.UpdateEstateType)

	return r
}
