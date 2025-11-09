package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"

	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
	"github.com/Parovozzzik/real-estate-portfolio/pkg/logging"
)

type TransactionTypeHandler struct {
	transactionTypeRepository *repositories.TransactionTypeRepository
}

func NewTransactionTypeHandler(transactionTypeRepository *repositories.TransactionTypeRepository) *TransactionTypeHandler {
	return &TransactionTypeHandler{transactionTypeRepository: transactionTypeRepository}
}

func (h *TransactionTypeHandler) GetTransactionTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := h.transactionTypeRepository.GetTransactionTypes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write(jsonData)
}

func (h *TransactionTypeHandler) CreateTransactionType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	createTransactionType := &models.CreateTransactionType{}
	err := json.NewDecoder(r.Body).Decode(createTransactionType)

	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newTransactionTypeId, err := h.transactionTypeRepository.CreateTransactionType(createTransactionType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	w.WriteHeader(http.StatusCreated)
	strTransactionTypeId := strconv.FormatInt(newTransactionTypeId, 10)
	fmt.Fprintf(w, `{"id": `+strTransactionTypeId+`}`)
}

func (h *TransactionTypeHandler) UpdateTransactionType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		panic(err)
	}

	updateTransactionType := &models.UpdateTransactionType{}
	updateTransactionType.Id = id
	err = json.NewDecoder(r.Body).Decode(updateTransactionType)

	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.transactionTypeRepository.UpdateTransactionType(updateTransactionType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	strTransactionTypeId := strconv.FormatInt(id, 10)
	fmt.Fprintf(w, `{"id": `+strTransactionTypeId+`}`)
}
