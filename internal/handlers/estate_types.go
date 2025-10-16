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

type EstateTypeHandler struct {
	estateTypeRepository *repositories.EstateTypeRepository
}

func NewEstateTypeHandler(estateTypeRepository *repositories.EstateTypeRepository) *EstateTypeHandler {
	return &EstateTypeHandler{estateTypeRepository: estateTypeRepository}
}

func (h *EstateTypeHandler) GetEstateTypes(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    jsonData, err := h.estateTypeRepository.GetEstateTypes()
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Write(jsonData)
}


func (h *EstateTypeHandler) CreateEstateType(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

	logging.Init()
    logger := logging.GetLogger()

	createEstateType := &models.CreateEstateType{}
	err := json.NewDecoder(r.Body).Decode(createEstateType)

	if err != nil {
	    logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

    newEstateTypeId, err := h.estateTypeRepository.CreateEstateType(createEstateType)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, `{"message": ` + err.Error() + `}`)
        return
    }

    w.WriteHeader(http.StatusCreated)
    strEstateTypeId := strconv.FormatInt(newEstateTypeId, 10)
    fmt.Fprintf(w, `{"id": ` + strEstateTypeId + `}`)
}

