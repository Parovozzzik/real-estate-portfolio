package handlers

import (
    "encoding/json"
    "net/http"
    "fmt"
    "strconv"

    "github.com/Parovozzzik/real-estate-portfolio/internal/logging"
    "github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
    "github.com/Parovozzzik/real-estate-portfolio/internal/models"
)

type EstateHandler struct {
	estateRepository *repositories.EstateRepository
}

func NewEstateHandler(estateRepository *repositories.EstateRepository) *EstateHandler {
	return &EstateHandler{estateRepository: estateRepository}
}

func (h *EstateHandler) GetEstates(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    jsonData, err := h.estateRepository.GetEstates()
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Write(jsonData)
}

func (h *EstateHandler) CreateEstate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

	logging.Init()
    logger := logging.GetLogger()

	createEstate := &models.CreateEstate{}
	err := json.NewDecoder(r.Body).Decode(createEstate)

	if err != nil {
	    logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

    newEstateId, err := h.estateRepository.CreateEstate(createEstate)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, `{"message": ` + err.Error() + `}`)
        return
    }

    w.WriteHeader(http.StatusCreated)
    strEstateId := strconv.FormatInt(newEstateId, 10)
    fmt.Fprintf(w, `{"id": ` + strEstateId + `}`)
}

