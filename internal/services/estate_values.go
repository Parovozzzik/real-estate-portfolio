package services

import (
	"encoding/json"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
	"github.com/Parovozzzik/real-estate-portfolio/pkg/logging"
	"net/http"
)

type EstateValuesService struct {
	estateValueRepository *repositories.EstateValueRepository
	transactionRepository *repositories.TransactionRepository
}

func NewEstateValuesService(
	estateValueRepository *repositories.EstateValueRepository,
	transactionRepository *repositories.TransactionRepository,
) *EstateValuesService {
	return &EstateValuesService{
		estateValueRepository: estateValueRepository,
		transactionRepository: transactionRepository,
	}
}

type RecalculateEstateValues struct {
	EstateId  int64  `json:"estate_id", db:"estate_id"`
	DateStart string `json:"date_start", db:"date_start"`
}

func (s *EstateValuesService) Recalculate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	logging.Init()
	logger := logging.GetLogger()
	recalculateEstateValues := &RecalculateEstateValues{}
	err := json.NewDecoder(r.Body).Decode(recalculateEstateValues)
	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transactions, err := s.transactionRepository.GetTransactionByEstateIdForValues(
		recalculateEstateValues.EstateId,
		recalculateEstateValues.DateStart)

	if err != nil {
		logger.Println(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = s.estateValueRepository.Upsert(&transactions)
	if err != nil {
		logger.Println(err.Error())

		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
