package models

import (
	"time"
)

type EstateValue struct {
	Id                int64     `json:"id"`
	EstateId          int64     `json:"estate_id"`
	Date              time.Time `json:"date"`
	Income            float64   `json:"income"`
	Expense           float64   `json:"expense"`
	Profit            float64   `json:"profit"`
	CumulativeIncome  float64   `json:"cumulative_income"`
	CumulativeExpense float64   `json:"cumulative_expense"`
	CumulativeProfit  float64   `json:"cumulative_profit"`
	Roi               float64   `json:"roi"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func NewEstateValue(id, estateId int64, date time.Time, income, expense, profit, cumulative_income, cumulative_expense, cumulative_profit, roi float64) *EstateValue {
	return &EstateValue{
		Id:                id,
		Date:              date,
		Income:            income,
		Expense:           expense,
		Profit:            profit,
		CumulativeIncome:  cumulative_income,
		CumulativeExpense: cumulative_expense,
		CumulativeProfit:  cumulative_profit,
		Roi:               roi,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

type CreateEstateValue struct {
	EstateId          int64     `json:"estate_id", db:"estate_id"`
	Date              time.Time `json:"date", db:"date"`
	Income            float64   `json:"income", db:"income"`
	Expense           float64   `json:"expense", db:"expense"`
	Profit            float64   `json:"profit", db:"profit"`
	CumulativeIncome  float64   `json:"cumulative_income", db:"cumulative_income"`
	CumulativeExpense float64   `json:"cumulative_expense", db:"cumulative_expense"`
	CumulativeProfit  float64   `json:"cumulative_profit", db:"cumulative_profit"`
	Roi               float64   `json:"roi", db:"roi"`
}

type UpdateEstateValue struct {
	Id                int64     `json:"id"`
	EstateId          int64     `json:"estate_id", db:"estate_id"`
	Date              time.Time `json:"date", db:"date"`
	Income            float64   `json:"income", db:"income"`
	Expense           float64   `json:"expense", db:"expense"`
	Profit            float64   `json:"profit", db:"profit"`
	CumulativeIncome  float64   `json:"cumulative_income", db:"cumulative_income"`
	CumulativeExpense float64   `json:"cumulative_expense", db:"cumulative_expense"`
	CumulativeProfit  float64   `json:"cumulative_profit", db:"cumulative_profit"`
	Roi               float64   `json:"roi", db:"roi"`
}

type FilterEstateValues struct {
	EstateId  *int64  `json:"estate_id"`
	Limit     *int    `json:"limit"`
	Page      *int64  `json:"page"`
	DateStart *string `json:"date_start"`
	DateEnd   *string `json:"date_end"`
	SortBy    *string `json:"sort_by"`
	SortOrder *string `json:"sort_order"`
}
