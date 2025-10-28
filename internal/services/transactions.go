package services

import (
	"encoding/json"
	"github.com/Parovozzzik/real-estate-portfolio/internal/logging"
	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
	"net/http"
	"time"
)

type TransactionService struct {
	transactionRepository              *repositories.TransactionRepository
	transactionGroupRepository         *repositories.TransactionGroupRepository
	transactionTypeRepository          *repositories.TransactionTypeRepository
	transactionFrequencyRepository     *repositories.TransactionFrequencyRepository
	transactionRepaymentPlanRepository *repositories.TransactionRepaymentPlanRepository
}

func NewTransactionService(
	transactionRepository *repositories.TransactionRepository,
	transactionGroupRepository *repositories.TransactionGroupRepository,
	transactionTypeRepository *repositories.TransactionTypeRepository,
	transactionFrequencyRepository *repositories.TransactionFrequencyRepository,
	transactionRepaymentPlanRepository *repositories.TransactionRepaymentPlanRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepository:              transactionRepository,
		transactionGroupRepository:         transactionGroupRepository,
		transactionTypeRepository:          transactionTypeRepository,
		transactionFrequencyRepository:     transactionFrequencyRepository,
		transactionRepaymentPlanRepository: transactionRepaymentPlanRepository,
	}
}

func (s *TransactionService) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	createTransactionGroupSettings := &models.CreateTransactionGroupSettings{}
	err := json.NewDecoder(r.Body).Decode(createTransactionGroupSettings)
	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := createTransactionGroupSettings
	if createTransactionGroupSettings.Direction == true && createTransactionGroupSettings.Regularity == true {
		result = s.RegularIncome(createTransactionGroupSettings)
	} else if createTransactionGroupSettings.Direction == true && createTransactionGroupSettings.Regularity == false {
		result = s.OneTimeIncome(createTransactionGroupSettings)
	} else if createTransactionGroupSettings.Direction == false && createTransactionGroupSettings.Regularity == true {
		result = s.RegularExpense(createTransactionGroupSettings)
	} else {
		result = s.OneTimeExpense(createTransactionGroupSettings)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (s *TransactionService) OneTimeIncome(createTransactionGroupSettings *models.CreateTransactionGroupSettings) *models.CreateTransactionGroupSettings {
	createTransactionGroup := &models.CreateTransactionGroup{}
	createTransactionGroup.EstateId = createTransactionGroupSettings.EstateId
	createTransactionGroup.Direction = createTransactionGroupSettings.Direction
	createTransactionGroup.Regularity = createTransactionGroupSettings.Regularity

	newTransactionGroupId, err := s.transactionGroupRepository.CreateTransactionGroup(createTransactionGroup)
	if err != nil {
		logging.Init()
		logger := logging.GetLogger()
		logger.Println(err.Error())
		return nil
	}

	createTransaction := &models.CreateTransaction{}
	createTransaction.GroupId = newTransactionGroupId
	createTransaction.Sum = createTransactionGroupSettings.Cost
	createTransaction.Date = time.Time(createTransactionGroupSettings.DateStart)
	createTransaction.TypeId = createTransactionGroupSettings.TypeId
	createTransaction.Comment = createTransactionGroupSettings.Comment

	_, err = s.transactionRepository.CreateTransaction(createTransaction)
	if err != nil {
		logging.Init()
		logger := logging.GetLogger()
		logger.Println(err.Error())
		return nil
	}

	return createTransactionGroupSettings
}

func (s *TransactionService) OneTimeExpense(createTransactionGroupSettings *models.CreateTransactionGroupSettings) *models.CreateTransactionGroupSettings {
	createTransactionGroup := &models.CreateTransactionGroup{}
	createTransactionGroup.EstateId = createTransactionGroupSettings.EstateId
	createTransactionGroup.Direction = createTransactionGroupSettings.Direction
	createTransactionGroup.Regularity = createTransactionGroupSettings.Regularity

	newTransactionGroupId, err := s.transactionGroupRepository.CreateTransactionGroup(createTransactionGroup)
	if err != nil {
		return nil
	}

	createTransaction := &models.CreateTransaction{}
	createTransaction.GroupId = newTransactionGroupId
	createTransaction.Sum = createTransactionGroupSettings.Cost
	createTransaction.Date = time.Time(createTransactionGroupSettings.DateStart)
	createTransaction.TypeId = createTransactionGroupSettings.TypeId
	createTransaction.Comment = createTransactionGroupSettings.Comment

	_, err = s.transactionRepository.CreateTransaction(createTransaction)
	if err != nil {
		return nil
	}

	return createTransactionGroupSettings
}

func (s *TransactionService) RegularIncome(createTransactionGroupSettings *models.CreateTransactionGroupSettings) *models.CreateTransactionGroupSettings {
	return &models.CreateTransactionGroupSettings{}
}

func (s *TransactionService) RegularExpense(createTransactionGroupSettings *models.CreateTransactionGroupSettings) *models.CreateTransactionGroupSettings {
	return &models.CreateTransactionGroupSettings{}
}

func (s *TransactionService) UpdateTransaction(w http.ResponseWriter, r *http.Request) {

}
