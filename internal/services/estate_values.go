package services

import (
	"encoding/json"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
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
	recalculateEstateValues := &RecalculateEstateValues{}
	err := json.NewDecoder(r.Body).Decode(recalculateEstateValues)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.RecalculateEstateValues(*recalculateEstateValues)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *EstateValuesService) RecalculateEstateValues(recalculateEstateValues RecalculateEstateValues) error {
	transactions, err := s.transactionRepository.GetTransactionByEstateIdForValues(
		recalculateEstateValues.EstateId,
		recalculateEstateValues.DateStart)
	if err != nil {
		return err
	}

	_, err = s.estateValueRepository.Upsert(&transactions)
	if err != nil {
		return err
	}

	return nil
}
