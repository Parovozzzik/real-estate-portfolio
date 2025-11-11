package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Parovozzzik/real-estate-portfolio/internal/database"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"

	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
	"github.com/Parovozzzik/real-estate-portfolio/pkg/logging"
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

	id, err := strconv.ParseInt(chi.URLParam(r, "transaction-id"), 10, 64)
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

	transaction, err := h.transactionRepository.GetTransactionById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	jsonData, err := json.Marshal(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func (h *TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	logging.Init()
	logger := logging.GetLogger()
	id, err := strconv.ParseInt(chi.URLParam(r, "transaction-id"), 10, 64)
	if err != nil {
		panic(err)
	}

	transaction, err := h.transactionRepository.GetTransactionById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	logger.Println(transaction)

	err = h.transactionRepository.Delete(id)
	if err != nil {
		logger.Println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	has, err := h.transactionRepository.HasTransactionsByGroupId(transaction.GroupId)
	logger.Println(has)

	if err != nil {
		logger.Println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	if has == false {
		db := database.GetDBInstance()
		transactionGroupRepository := repositories.NewTransactionGroupRepository(db)
		err = transactionGroupRepository.DeleteEmptyTransactionGroup(transaction.GroupId)
		if err != nil {
			logger.Println(err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
