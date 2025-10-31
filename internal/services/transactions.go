package services

import (
	"encoding/json"
	"github.com/Parovozzzik/real-estate-portfolio/internal/logging"
	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
	"github.com/Parovozzzik/real-estate-portfolio/internal/utils"
	"math"
	"net/http"
	"time"
)

type TransactionService struct {
	transactionRepository              *repositories.TransactionRepository
	transactionGroupRepository         *repositories.TransactionGroupRepository
	transactionTypeRepository          *repositories.TransactionTypeRepository
	transactionFrequencyRepository     *repositories.TransactionFrequencyRepository
	transactionRepaymentPlanRepository *repositories.TransactionRepaymentPlanRepository
	transactionGroupSettingRepository  *repositories.TransactionGroupSettingRepository
}

func NewTransactionService(
	transactionRepository *repositories.TransactionRepository,
	transactionGroupRepository *repositories.TransactionGroupRepository,
	transactionTypeRepository *repositories.TransactionTypeRepository,
	transactionFrequencyRepository *repositories.TransactionFrequencyRepository,
	transactionRepaymentPlanRepository *repositories.TransactionRepaymentPlanRepository,
	transactionGroupSettingRepository *repositories.TransactionGroupSettingRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepository:              transactionRepository,
		transactionGroupRepository:         transactionGroupRepository,
		transactionTypeRepository:          transactionTypeRepository,
		transactionFrequencyRepository:     transactionFrequencyRepository,
		transactionRepaymentPlanRepository: transactionRepaymentPlanRepository,
		transactionGroupSettingRepository:  transactionGroupSettingRepository,
	}
}

type CreateFullTransactionGroup struct {
	EstateId        int64            `json:"estate_id", db:"estate_id"`
	TypeId          int64            `json:"type_id", db:"type_id"`
	Direction       bool             `json:"direction", db:"direction"`
	Regularity      bool             `json:"regularity", db:"regularity"`
	Name            string           `json:"name", db:"name"`
	Comment         *string          `json:"comment", db:"comment"`
	Cost            float64          `json:"cost", db:"cost"`
	DownPayment     float64          `json:"down_payment", db:"down_payment"`
	OwnFunds        float64          `json:"own_funds", db:"own_funds"`
	ThirdPartyFunds float64          `json:"third_party_funds", db:"third_party_funds"`
	InterestRate    float64          `json:"interest_rate", db:"interest_rate"`
	FrequencyId     int64            `json:"frequency_id", db:"frequency_id"`
	RepaymentPlanId int64            `json:"repayment_plan_id", db:"repayment_plan_id"`
	DateStart       utils.CustomTime `json:"date_start", db:"date_start"`
	LoanTerm        int              `json:"loan_term", db:"loan_term"`
	Payday          int              `json:"payday", db:"payday"`
	PaydayOnWorkday bool             `json:"payday_on_workday", db:"payday_on_workday"`
}

type UpdateTransactionGroupSettings struct {
	Id              int64            `json:"id", db:"id"`
	TypeId          int64            `json:"type_id", db:"type_id"`
	Name            string           `json:"name", db:"name"`
	Comment         string           `json:"comment", db:"comment"`
	Cost            float64          `json:"cost", db:"cost"`
	DownPayment     float64          `json:"down_payment", db:"down_payment"`
	OwnFunds        float64          `json:"own_funds", db:"own_funds"`
	ThirdPartyFunds float64          `json:"third_party_funds", db:"third_party_funds"`
	InterestRate    float64          `json:"interest_rate", db:"interest_rate"`
	FrequencyId     int64            `json:"frequency_id", db:"frequency_id"`
	RepaymentPlanId int64            `json:"repayment_plan_id", db:"repayment_plan_id"`
	DateStart       utils.CustomTime `json:"date_start", db:"date_start"`
	LoanTerm        int              `json:"loan_term", db:"loan_term"`
	Payday          int              `json:"payday", db:"payday"`
	PaydayOnWorkday bool             `json:"payday_on_workday", db:"payday_on_workday"`
	EstateId        int64            `json:"estate_id", db:"estate_id"`
	Direction       bool             `json:"direction", db:"direction"`
	Regularity      bool             `json:"regularity", db:"regularity"`
}

func (s *TransactionService) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	createTransactionGroupSettings := &CreateFullTransactionGroup{}
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

func (s *TransactionService) OneTimeIncome(createTransactionGroupSettings *CreateFullTransactionGroup) *CreateFullTransactionGroup {
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

func (s *TransactionService) OneTimeExpense(createTransactionGroupSettings *CreateFullTransactionGroup) *CreateFullTransactionGroup {
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

func (s *TransactionService) RegularIncome(createTransactionGroupSettings *CreateFullTransactionGroup) *CreateFullTransactionGroup {
	newTransactionGroupSetting := &models.CreateTransactionGroupSetting{}
	newTransactionGroupSetting.Name = createTransactionGroupSettings.Name
	newTransactionGroupSetting.Cost = createTransactionGroupSettings.Cost
	newTransactionGroupSetting.InterestRate = createTransactionGroupSettings.InterestRate
	newTransactionGroupSetting.FrequencyId = createTransactionGroupSettings.FrequencyId
	newTransactionGroupSetting.DateStart = time.Time(createTransactionGroupSettings.DateStart)
	newTransactionGroupSetting.LoanTerm = createTransactionGroupSettings.LoanTerm
	newTransactionGroupSetting.Payday = createTransactionGroupSettings.Payday
	newTransactionGroupSetting.PaydayOnWorkday = createTransactionGroupSettings.PaydayOnWorkday
	newTransactionGroupSettingsId, err := s.transactionGroupSettingRepository.CreateTransactionGroupSetting(newTransactionGroupSetting)
	if err != nil {
		logging.Init()
		logger := logging.GetLogger()
		logger.Println(err.Error())
		return nil
	}

	newTransactionGroup := &models.CreateTransactionGroup{}
	newTransactionGroup.EstateId = createTransactionGroupSettings.EstateId
	newTransactionGroup.SettingId = &newTransactionGroupSettingsId
	newTransactionGroup.Direction = createTransactionGroupSettings.Direction
	newTransactionGroup.Regularity = createTransactionGroupSettings.Regularity

	newTransactionGroupId, err := s.transactionGroupRepository.CreateTransactionGroup(newTransactionGroup)
	if err != nil {
		logging.Init()
		logger := logging.GetLogger()
		logger.Println(err.Error())
		return nil
	}

	dates := s.GetPaymentDates(
		newTransactionGroupSetting.DateStart,
		newTransactionGroupSetting.LoanTerm,
		newTransactionGroupSetting.FrequencyId)

	for _, date := range dates {
		newTransaction := &models.CreateTransaction{}
		newTransaction.GroupId = newTransactionGroupId
		newTransaction.Sum = createTransactionGroupSettings.Cost
		newTransaction.Date = date
		newTransaction.TypeId = createTransactionGroupSettings.TypeId
		newTransaction.Comment = createTransactionGroupSettings.Comment

		_, err = s.transactionRepository.CreateTransaction(newTransaction)
		if err != nil {
			logging.Init()
			logger := logging.GetLogger()
			logger.Println(err.Error())
			return nil
		}
	}

	return createTransactionGroupSettings
}

func (s *TransactionService) RegularExpense(createTransactionGroupSettings *CreateFullTransactionGroup) *CreateFullTransactionGroup {
	principal := createTransactionGroupSettings.Cost
	annualInterestRate := createTransactionGroupSettings.InterestRate
	loanTermMonths := createTransactionGroupSettings.LoanTerm

	monthlyPayment, _ := calculateLoan(principal, annualInterestRate, loanTermMonths)

	newTransactionGroupSetting := &models.CreateTransactionGroupSetting{}
	newTransactionGroupSetting.Name = createTransactionGroupSettings.Name
	newTransactionGroupSetting.Cost = createTransactionGroupSettings.Cost
	newTransactionGroupSetting.FrequencyId = createTransactionGroupSettings.FrequencyId
	newTransactionGroupSetting.DateStart = time.Time(createTransactionGroupSettings.DateStart)
	newTransactionGroupSetting.LoanTerm = createTransactionGroupSettings.LoanTerm
	newTransactionGroupSetting.Payday = createTransactionGroupSettings.Payday
	newTransactionGroupSetting.PaydayOnWorkday = createTransactionGroupSettings.PaydayOnWorkday
	newTransactionGroupSettingsId, err := s.transactionGroupSettingRepository.CreateTransactionGroupSetting(newTransactionGroupSetting)
	if err != nil {
		return nil
	}

	newTransactionGroup := &models.CreateTransactionGroup{}
	newTransactionGroup.EstateId = createTransactionGroupSettings.EstateId
	newTransactionGroup.SettingId = &newTransactionGroupSettingsId
	newTransactionGroup.Direction = createTransactionGroupSettings.Direction
	newTransactionGroup.Regularity = createTransactionGroupSettings.Regularity

	newTransactionGroupId, err := s.transactionGroupRepository.CreateTransactionGroup(newTransactionGroup)
	if err != nil {
		logging.Init()
		logger := logging.GetLogger()
		logger.Println(err.Error())
		return nil
	}

	dates := s.GetPaymentDates(
		newTransactionGroupSetting.DateStart,
		newTransactionGroupSetting.LoanTerm,
		newTransactionGroupSetting.FrequencyId)

	for _, date := range dates {
		newTransaction := &models.CreateTransaction{}
		newTransaction.GroupId = newTransactionGroupId
		newTransaction.Sum = monthlyPayment
		newTransaction.Date = date
		newTransaction.TypeId = createTransactionGroupSettings.TypeId
		newTransaction.Comment = createTransactionGroupSettings.Comment

		_, err = s.transactionRepository.CreateTransaction(newTransaction)
		if err != nil {
			logging.Init()
			logger := logging.GetLogger()
			logger.Println(err.Error())
			return nil
		}
	}

	return createTransactionGroupSettings
}

func (s *TransactionService) UpdateTransaction(w http.ResponseWriter, r *http.Request) {

}

func (s *TransactionService) GetPaymentDates(dateStart time.Time, loanTerm int, frequencyId int64) []time.Time {
	dates := make([]time.Time, loanTerm)
	dates[0] = dateStart
	if frequencyId == 3 { // month
		for i := 1; i < loanTerm; i++ {
			dates[i] = dateStart.AddDate(0, i, 0)
		}
	}

	logging.Init()
	logger := logging.GetLogger()
	logger.Println(dates)

	return dates
}

func calculateLoan(principal float64, annualInterestRate float64, loanTermMonths int) (monthlyPayment float64, totalInterest float64) {
	monthlyInterestRate := annualInterestRate / 100 / 12

	// P = S * [r* (1+r)^n] / [(1+r)^n – 1]
	power := math.Pow(1+monthlyInterestRate, float64(loanTermMonths))
	monthlyPayment = principal * (monthlyInterestRate * power) / (power - 1)

	// Переплата = (Ежемесячный платеж * срок) - Сумма кредита
	totalRepayment := monthlyPayment*float64(loanTermMonths) - principal

	return monthlyPayment, totalRepayment
}
