package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
)

type TransactionTypeRepository struct {
	db *sql.DB
}

func NewTransactionTypeRepository(db *sql.DB) *TransactionTypeRepository {
	return &TransactionTypeRepository{db: db}
}

func (u *TransactionTypeRepository) GetTransactionTypes() ([]byte, error) {
	rows, err := u.db.Query(
		"SELECT id, name FROM real_estate_portfolio.rep_transaction_types")
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

func (u *TransactionTypeRepository) CreateTransactionType(createTransactionType *models.CreateTransactionType) (int64, error) {
	result, err := u.db.Exec(
		"INSERT INTO real_estate_portfolio.rep_transaction_types (name) VALUES (?)", createTransactionType.Name)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (u *TransactionTypeRepository) UpdateTransactionType(updateTransactionType *models.UpdateTransactionType) error {
	_, err := u.db.Exec(
		"UPDATE real_estate_portfolio.rep_transaction_types SET name = ? WHERE id = ?",
		updateTransactionType.Name,
		updateTransactionType.Id)
	return err
}
