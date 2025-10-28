package models

import (
	"time"
)

type TransactionGroupSetting struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	Cost            float64   `json:"cost"`
	DownPayment     float64   `json:"down_payment"`
	OwnFunds        float64   `json:"own_funds"`
	ThirdPartyFunds float64   `json:"third_party_funds"`
	InterestRate    float64   `json:"interest_rate"`
	FrequencyId     int64     `json:"frequency_id"`
	RepaymentPlanId int64     `json:"repayment_plan_id"`
	DateStart       time.Time `json:"date_start"`
	LoanTerm        int       `json:"loan_term"`
	Payday          int       `json:"payday"`
	PaydayOnWorkday bool      `json:"payday_on_workday"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateTransactionGroupSettings struct {
	EstateId        int64      `json:"estate_id", db:"estate_id"`
	TypeId          int64      `json:"type_id", db:"type_id"`
	Direction       bool       `json:"direction", db:"direction"`
	Regularity      bool       `json:"regularity", db:"regularity"`
	Name            string     `json:"name", db:"name"`
	Comment         *string    `json:"comment", db:"comment"`
	Cost            float64    `json:"cost", db:"cost"`
	DownPayment     float64    `json:"down_payment", db:"down_payment"`
	OwnFunds        float64    `json:"own_funds", db:"own_funds"`
	ThirdPartyFunds float64    `json:"third_party_funds", db:"third_party_funds"`
	InterestRate    float64    `json:"interest_rate", db:"interest_rate"`
	FrequencyId     int64      `json:"frequency_id", db:"frequency_id"`
	RepaymentPlanId int64      `json:"repayment_plan_id", db:"repayment_plan_id"`
	DateStart       CustomTime `json:"date_start", db:"date_start"`
	LoanTerm        int        `json:"loan_term", db:"loan_term"`
	Payday          int        `json:"payday", db:"payday"`
	PaydayOnWorkday bool       `json:"payday_on_workday", db:"payday_on_workday"`
}

type UpdateTransactionGroupSettings struct {
	Id              int64      `json:"id", db:"id"`
	TypeId          int64      `json:"type_id", db:"type_id"`
	Name            string     `json:"name", db:"name"`
	Comment         string     `json:"comment", db:"comment"`
	Cost            float64    `json:"cost", db:"cost"`
	DownPayment     float64    `json:"down_payment", db:"down_payment"`
	OwnFunds        float64    `json:"own_funds", db:"own_funds"`
	ThirdPartyFunds float64    `json:"third_party_funds", db:"third_party_funds"`
	InterestRate    float64    `json:"interest_rate", db:"interest_rate"`
	FrequencyId     int64      `json:"frequency_id", db:"frequency_id"`
	RepaymentPlanId int64      `json:"repayment_plan_id", db:"repayment_plan_id"`
	DateStart       CustomTime `json:"date_start", db:"date_start"`
	LoanTerm        int        `json:"loan_term", db:"loan_term"`
	Payday          int        `json:"payday", db:"payday"`
	PaydayOnWorkday bool       `json:"payday_on_workday", db:"payday_on_workday"`
	EstateId        int64      `json:"estate_id", db:"estate_id"`
	Direction       bool       `json:"direction", db:"direction"`
	Regularity      bool       `json:"regularity", db:"regularity"`
}
