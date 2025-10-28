package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"

	"github.com/Parovozzzik/real-estate-portfolio/internal/logging"
	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
)

type TransactionFrequencyHandler struct {
	transactionTypeRepository *repositories.TransactionFrequencyRepository
}

func NewTransactionFrequencyHandler(transactionTypeRepository *repositories.TransactionFrequencyRepository) *TransactionFrequencyHandler {
	return &TransactionFrequencyHandler{transactionTypeRepository: transactionTypeRepository}
}

func (h *TransactionFrequencyHandler) GetTransactionFrequencies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := h.transactionTypeRepository.GetTransactionFrequencies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write(jsonData)
}

func (h *TransactionFrequencyHandler) CreateTransactionFrequency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	createTransactionFrequency := &models.CreateTransactionFrequency{}
	err := json.NewDecoder(r.Body).Decode(createTransactionFrequency)

	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newTransactionFrequencyId, err := h.transactionTypeRepository.CreateTransactionFrequency(createTransactionFrequency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	w.WriteHeader(http.StatusCreated)
	strTransactionFrequencyId := strconv.FormatInt(newTransactionFrequencyId, 10)
	fmt.Fprintf(w, `{"id": `+strTransactionFrequencyId+`}`)
}

func (h *TransactionFrequencyHandler) UpdateTransactionFrequency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		panic(err)
	}

	updateTransactionFrequency := &models.UpdateTransactionFrequency{}
	updateTransactionFrequency.Id = id
	err = json.NewDecoder(r.Body).Decode(updateTransactionFrequency)

	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.transactionTypeRepository.UpdateTransactionFrequency(updateTransactionFrequency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	strTransactionFrequencyId := strconv.FormatInt(id, 10)
	fmt.Fprintf(w, `{"id": `+strTransactionFrequencyId+`}`)
}
