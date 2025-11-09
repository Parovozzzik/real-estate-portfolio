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

type TransactionGroupHandler struct {
	transactionGroupRepository *repositories.TransactionGroupRepository
}

func NewTransactionGroupHandler(transactionGroupRepository *repositories.TransactionGroupRepository) *TransactionGroupHandler {
	return &TransactionGroupHandler{transactionGroupRepository: transactionGroupRepository}
}

func (h *TransactionGroupHandler) GetTransactionGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := h.transactionGroupRepository.GetTransactionGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write(jsonData)
}

func (h *TransactionGroupHandler) CreateTransactionGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	createTransactionGroup := &models.CreateTransactionGroup{}
	err := json.NewDecoder(r.Body).Decode(createTransactionGroup)

	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newTransactionGroupId, err := h.transactionGroupRepository.CreateTransactionGroup(createTransactionGroup)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	w.WriteHeader(http.StatusCreated)
	strTransactionGroupId := strconv.FormatInt(newTransactionGroupId, 10)
	fmt.Fprintf(w, `{"id": `+strTransactionGroupId+`}`)
}

func (h *TransactionGroupHandler) UpdateTransactionGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		panic(err)
	}

	updateTransactionGroup := &models.UpdateTransactionGroup{}
	updateTransactionGroup.Id = id
	err = json.NewDecoder(r.Body).Decode(updateTransactionGroup)

	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.transactionGroupRepository.UpdateTransactionGroup(updateTransactionGroup)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	strTransactionGroupId := strconv.FormatInt(id, 10)
	fmt.Fprintf(w, `{"id": `+strTransactionGroupId+`}`)
}
