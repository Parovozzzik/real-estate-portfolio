package repositories

import (
	"database/sql"

	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
)

type TransactionGroupSettingRepository struct {
	db *sql.DB
}

func NewTransactionGroupSettingRepository(db *sql.DB) *TransactionGroupSettingRepository {
	return &TransactionGroupSettingRepository{db: db}
}

func (u TransactionGroupSettingRepository) CreateTransactionGroupSetting(createTransactionGroupSetting *models.CreateTransactionGroupSetting) (int64, error) {
	result, err := u.db.Exec(
		"INSERT INTO real_estate_portfolio.rep_transaction_group_settings (name, cost, down_payment, own_funds, third_party_funds, interest_rate, frequency_id, repayment_plan_id, date_start, loan_term, payday, payday_on_workday) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		createTransactionGroupSetting.Name,
		createTransactionGroupSetting.Cost,
		createTransactionGroupSetting.DownPayment,
		createTransactionGroupSetting.OwnFunds,
		createTransactionGroupSetting.ThirdPartyFunds,
		createTransactionGroupSetting.InterestRate,
		createTransactionGroupSetting.FrequencyId,
		createTransactionGroupSetting.RepaymentPlanId,
		createTransactionGroupSetting.DateStart,
		createTransactionGroupSetting.LoanTerm,
		createTransactionGroupSetting.Payday,
		createTransactionGroupSetting.PaydayOnWorkday)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}
