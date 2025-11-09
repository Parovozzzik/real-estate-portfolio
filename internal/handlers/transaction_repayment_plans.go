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

type TransactionRepaymentPlanHandler struct {
	transactionTypeRepository *repositories.TransactionRepaymentPlanRepository
}

func NewTransactionRepaymentPlanHandler(transactionTypeRepository *repositories.TransactionRepaymentPlanRepository) *TransactionRepaymentPlanHandler {
	return &TransactionRepaymentPlanHandler{transactionTypeRepository: transactionTypeRepository}
}

func (h *TransactionRepaymentPlanHandler) GetTransactionRepaymentPlans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := h.transactionTypeRepository.GetTransactionRepaymentPlans()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write(jsonData)
}

func (h *TransactionRepaymentPlanHandler) CreateTransactionRepaymentPlan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	createTransactionRepaymentPlan := &models.CreateTransactionRepaymentPlan{}
	err := json.NewDecoder(r.Body).Decode(createTransactionRepaymentPlan)

	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newTransactionRepaymentPlanId, err := h.transactionTypeRepository.CreateTransactionRepaymentPlan(createTransactionRepaymentPlan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	w.WriteHeader(http.StatusCreated)
	strTransactionRepaymentPlanId := strconv.FormatInt(newTransactionRepaymentPlanId, 10)
	fmt.Fprintf(w, `{"id": `+strTransactionRepaymentPlanId+`}`)
}

func (h *TransactionRepaymentPlanHandler) UpdateTransactionRepaymentPlan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		panic(err)
	}

	updateTransactionRepaymentPlan := &models.UpdateTransactionRepaymentPlan{}
	updateTransactionRepaymentPlan.Id = id
	err = json.NewDecoder(r.Body).Decode(updateTransactionRepaymentPlan)

	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.transactionTypeRepository.UpdateTransactionRepaymentPlan(updateTransactionRepaymentPlan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	strTransactionRepaymentPlanId := strconv.FormatInt(id, 10)
	fmt.Fprintf(w, `{"id": `+strTransactionRepaymentPlanId+`}`)
}
