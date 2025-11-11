package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
	"github.com/Parovozzzik/real-estate-portfolio/pkg/logging"
)

type TransactionGroupRepository struct {
	db *sql.DB
}

func NewTransactionGroupRepository(db *sql.DB) *TransactionGroupRepository {
	return &TransactionGroupRepository{db: db}
}

func (u *TransactionGroupRepository) GetTransactionGroups() ([]byte, error) {
	rows, err := u.db.Query(
		"SELECT id, estate_id, direction, regularity FROM real_estate_portfolio.rep_transaction_groups")
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

func (u *TransactionGroupRepository) CreateTransactionGroup(createTransactionGroup *models.CreateTransactionGroup) (int64, error) {
	result, err := u.db.Exec(
		"INSERT INTO real_estate_portfolio.rep_transaction_groups (estate_id, setting_id, direction, regularity) VALUES (?, ?, ?, ?)",
		createTransactionGroup.EstateId,
		createTransactionGroup.SettingId,
		createTransactionGroup.Direction,
		createTransactionGroup.Regularity)
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

func (u *TransactionGroupRepository) UpdateTransactionGroup(updateTransactionGroup *models.UpdateTransactionGroup) error {
	_, err := u.db.Exec(
		"UPDATE real_estate_portfolio.rep_transaction_groups SET estate_id = ?, setting_id = ?, direction = ?, regularity = ? WHERE id = ?",
		updateTransactionGroup.EstateId,
		updateTransactionGroup.SettingId,
		updateTransactionGroup.Direction,
		updateTransactionGroup.Regularity,
		updateTransactionGroup.Id)
	return err
}

func (u *TransactionGroupRepository) DeleteEmptyTransactionGroup(id int64) error {
	_, err := u.db.Exec("DELETE FROM real_estate_portfolio.rep_transaction_groups WHERE id = ?", id)

	return err
}

func (u *TransactionGroupRepository) GetTransactionGroupsByEstateId(estateId int64) ([]map[string]interface{}, error) {
	rows, err := u.db.Query(
		"SELECT id, setting_id FROM real_estate_portfolio.rep_transaction_groups WHERE estate_id = ?", estateId)
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

	return tableData, nil
}

func (u TransactionGroupRepository) DeleteByIds(ids []int64) error {
	placeholders := make([]string, len(ids))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	inClause := strings.Join(placeholders, ",")
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	if len(args) > 0 {
		_, err := u.db.Exec("DELETE FROM real_estate_portfolio.rep_transaction_groups WHERE id IN ("+inClause+")", args...)
		return err
	}

	return nil
}
