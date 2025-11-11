package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
)

type EstateValueRepository struct {
	db *sql.DB
}

func NewEstateValueRepository(db *sql.DB) *EstateValueRepository {
	return &EstateValueRepository{db: db}
}

func (u *EstateValueRepository) GetEstateValues() ([]byte, error) {
	rows, err := u.db.Query(
		"SELECT id, name, description, estate_type_id, user_id FROM real_estate_portfolio.rep_estates")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	if err != nil {
		fmt.Println("Error...")
	}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("Error...")
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		fmt.Println("Error...")
	}

	return jsonData, nil
}

func (u *EstateValueRepository) CreateEstateValue(createEstateValue *models.CreateEstateValue) (int64, error) {
	result, err := u.db.Exec(
		"INSERT INTO real_estate_portfolio.rep_estate_values (estate_id, date, income, expense, profit, cumulative_income, cumulative_expense, cumulative_profit, roi) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		createEstateValue.EstateId,
		createEstateValue.Date,
		createEstateValue.Income,
		createEstateValue.Expense,
		createEstateValue.Profit,
		createEstateValue.CumulativeIncome,
		createEstateValue.CumulativeExpense,
		createEstateValue.CumulativeProfit,
		createEstateValue.Roi)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}
func (u *EstateValueRepository) UpdateEstateValue(updateEstateValue *models.UpdateEstateValue) (int64, error) {
	_, err := u.db.Exec(
		"UPDATE real_estate_portfolio.rep_estate_values SET estate_id = ?, date = ?, income = ?, expense = ?, profit = ?, cumulative_income = ?, cumulative_expense = ?, cumulative_profit = ?, roi = ? WHERE id = ?",
		updateEstateValue.EstateId,
		updateEstateValue.Date,
		updateEstateValue.Income,
		updateEstateValue.Expense,
		updateEstateValue.Profit,
		updateEstateValue.CumulativeIncome,
		updateEstateValue.CumulativeExpense,
		updateEstateValue.CumulativeProfit,
		updateEstateValue.Roi,
		updateEstateValue.Id)
	if err != nil {
		return 0, err
	}

	return updateEstateValue.Id, nil
}

func (u *EstateValueRepository) Upsert(transactions *[]map[string]interface{}) (int64, error) {
	query := "INSERT IGNORE INTO real_estate_portfolio.rep_estate_values " +
		"(estate_id, date, income, expense, profit, " +
		"cumulative_income, cumulative_expense, cumulative_profit, roi) " +
		"VALUES "

	params := []any{}
	for _, transaction := range *transactions {
		query += "(?, ?, ?, ?, ?, ?, ?, ?, ?), "

		date := fmt.Sprintf("%d-%02d-15",
			transaction["year"].(int64),
			transaction["month"].(int64))

		params = append(params,
			transaction["estate_id"],
			date,
			transaction["income"],
			transaction["expense"],
			transaction["profit"],
			transaction["cumulative_income"],
			transaction["cumulative_expense"],
			transaction["cumulative_profit"],
			transaction["roi"])
	}
	query = strings.Trim(query, ", ")
	query += "ON DUPLICATE KEY UPDATE " +
		"estate_id = VALUES(estate_id), " +
		"date = VALUES(date), " +
		"income = VALUES(income), " +
		"expense = VALUES(expense), " +
		"profit = VALUES(profit), " +
		"cumulative_income = VALUES(cumulative_income), " +
		"cumulative_expense = VALUES(cumulative_expense), " +
		"cumulative_profit = VALUES(cumulative_profit), " +
		"roi = VALUES(roi)"
	_, err := u.db.Exec(query, params...)

	return 0, err
}

func (u EstateValueRepository) DeleteByEstateId(estateId int64) error {
	_, err := u.db.Exec("DELETE FROM real_estate_portfolio.rep_estate_values WHERE estate_id = ?", estateId)

	return err
}
