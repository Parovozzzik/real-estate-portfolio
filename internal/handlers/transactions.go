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

type TransactionHandler struct {
	transactionRepository *repositories.TransactionRepository
}

func NewTransactionHandler(transactionRepository *repositories.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{transactionRepository: transactionRepository}
}

func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := h.transactionRepository.GetTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write(jsonData)
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	createTransaction := &models.CreateTransaction{}
	err := json.NewDecoder(r.Body).Decode(createTransaction)

	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newTransactionId, err := h.transactionRepository.CreateTransaction(createTransaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	w.WriteHeader(http.StatusCreated)
	strTransactionId := strconv.FormatInt(newTransactionId, 10)
	fmt.Fprintf(w, `{"id": `+strTransactionId+`}`)
}

func (h *TransactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		panic(err)
	}

	updateTransaction := &models.UpdateTransaction{}
	updateTransaction.Id = id
	err = json.NewDecoder(r.Body).Decode(updateTransaction)

	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.transactionRepository.UpdateTransaction(updateTransaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	strTransactionId := strconv.FormatInt(id, 10)
	fmt.Fprintf(w, `{"id": `+strTransactionId+`}`)
}
