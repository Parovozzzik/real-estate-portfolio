package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Parovozzzik/real-estate-portfolio/internal/logging"
	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (u *TransactionRepository) GetTransactions() ([]byte, error) {
	rows, err := u.db.Query(
		"SELECT id, sum, date FROM real_estate_portfolio.rep_transaction")
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

func (u *TransactionRepository) CreateTransaction(createTransaction *models.CreateTransaction) (int64, error) {
	result, err := u.db.Exec(
		"INSERT INTO real_estate_portfolio.rep_transactions (group_id, type_id, sum, date, comment) VALUES (?, ?, ?, ?, ?)",
		createTransaction.GroupId,
		createTransaction.TypeId,
		createTransaction.Sum,
		createTransaction.Date,
		createTransaction.Comment,
	)
	if err != nil {
		logging.Init()
		logger := logging.GetLogger()
		logger.Println(err.Error())
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		logging.Init()
		logger := logging.GetLogger()
		logger.Println(err.Error())
		return 0, err
	}

	return lastInsertID, nil
}

func (u *TransactionRepository) UpdateTransaction(updateTransaction *models.UpdateTransaction) error {
	_, err := u.db.Exec(
		"UPDATE real_estate_portfolio.rep_transactions SET group_id = ?, type_id = ?, sum = ?, date = ?, comment = ? WHERE id = ?",
		updateTransaction.GroupId,
		updateTransaction.TypeId,
		updateTransaction.Sum,
		updateTransaction.Date,
		updateTransaction.Comment,
		updateTransaction.Id)
	return err
}
