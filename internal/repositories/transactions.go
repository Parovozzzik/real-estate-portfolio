package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
	"github.com/Parovozzzik/real-estate-portfolio/pkg/logging"
	"log"
	"strings"
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
		return 0, err
	}

	return lastInsertID, nil
}

func (u *TransactionRepository) UpdateTransaction(updateTransaction *models.UpdateTransaction) error {
	params := []any{}
	query := "UPDATE real_estate_portfolio.rep_transactions SET "
	if updateTransaction != nil {
		if updateTransaction.GroupId != nil {
			query += "group_id = ?, "
			params = append(params, updateTransaction.GroupId)
		}
		if updateTransaction.TypeId != nil {
			query += "type_id = ?, "
			params = append(params, updateTransaction.TypeId)
		}
		if updateTransaction.Sum != nil {
			query += "sum = ?, "
			params = append(params, updateTransaction.Sum)
		}
		if updateTransaction.Date != nil {
			query += "date = ?, "
			params = append(params, updateTransaction.Date)
		}
		if updateTransaction.Comment != nil {
			query += "comment = ?, "
			params = append(params, updateTransaction.Comment)
		}
	}

	if len(params) > 0 {
		query = strings.Trim(query, ", ")
		query += " WHERE id = ?"
		params = append(params, updateTransaction.Id)
	}

	_, err := u.db.Exec(query, params...)

	return err
}

func (u *TransactionRepository) GetTransactionById(id int64) (*models.FullTransaction, error) {
	transactionData, err := u.db.Query(
		"SELECT tr.id, tr.sum, tr.date, tr.comment, tr.type_id, trt.name AS type_name, tr.group_id "+
			"FROM real_estate_portfolio.rep_transactions tr "+
			"JOIN real_estate_portfolio.rep_transaction_types trt ON trt.id = tr.type_id "+
			"WHERE tr.id = ?",
		id)
	if err != nil {
		return nil, err
	}
	defer transactionData.Close()

	transaction := &models.FullTransaction{}
	transactionData.Next()
	err = transactionData.Scan(&transaction.Id, &transaction.Sum, &transaction.Date, &transaction.Comment, &transaction.TypeId, &transaction.TypeName, &transaction.GroupId)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (u *TransactionRepository) HasTransactionsByGroupId(groupId int64) (bool, error) {
	var rowsCount int
	countQuery := "SELECT COUNT(tr.id) " +
		"FROM real_estate_portfolio.rep_transactions tr " +
		"JOIN real_estate_portfolio.rep_transaction_groups trg ON trg.id = tr.group_id " +
		"WHERE trg.id = ?"
	err := u.db.QueryRow(countQuery, groupId).Scan(&rowsCount)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return rowsCount > 0, nil
}

func (u *TransactionRepository) Delete(id int64) error {
	_, err := u.db.Exec("DELETE FROM real_estate_portfolio.rep_transactions WHERE id = ?", id)

	return err
}

func (u TransactionRepository) DeleteByGroupIds(ids []int64) error {
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
		_, err := u.db.Exec("DELETE FROM real_estate_portfolio.rep_transactions WHERE group_id IN ("+inClause+")", args...)
		return err
	}

	return nil
}

func (u *TransactionRepository) GetTransactionByEstateIdForValues(estateId int64, dateStart string) ([]map[string]interface{}, error) {
	params := []any{}
	query :=
		`SELECT 
       estate_id,
       year,
       month,
       income,
       expense,
       profit,
       cumulative_income,
       cumulative_expense,
       cumulative_profit,
       (cumulative_income - (cumulative_expense * -1)) / (cumulative_expense * -1) * 100 as roi
		FROM (SELECT rtg.estate_id,
			 YEAR(tr.date)                                                  as year,
             MONTH(tr.date)                                                 as month,
             SUM(IF(rtt.direction = 1, tr.sum, 0))                          as income,
             SUM(IF(rtt.direction = 1, 0, tr.sum * -1))                     as expense,
             SUM(IF(rtt.direction = 1, tr.sum, tr.sum * -1))                as profit,
             (SELECT SUM(IF(rtt2.direction = 1, tr2.sum, 0)) as income
              FROM rep_transactions tr2
              JOIN rep_transaction_types rtt2 on rtt2.id = tr2.type_id
              JOIN rep_transaction_groups rtg2 on tr2.group_id = rtg2.id
              WHERE rtg2.estate_id = ? AND (MONTH(tr2.date) + YEAR(tr2.date) * 12) <= (month + year * 12)) as cumulative_income,
             (SELECT SUM(IF(rtt2.direction = 1, 0, tr2.sum * -1)) as expense
              FROM rep_transactions tr2
               JOIN rep_transaction_types rtt2 on rtt2.id = tr2.type_id
               JOIN rep_transaction_groups rtg2 on tr2.group_id = rtg2.id
              WHERE rtg2.estate_id = ? AND (MONTH(tr2.date) + YEAR(tr2.date) * 12) <= (month + year * 12)) as cumulative_expense,
             (SELECT SUM(IF(rtt2.direction = 1, tr2.sum, tr2.sum * -1)) as profit
              FROM rep_transactions tr2
               JOIN rep_transaction_types rtt2 on rtt2.id = tr2.type_id
               JOIN rep_transaction_groups rtg2 on tr2.group_id = rtg2.id
              WHERE rtg2.estate_id = ? AND (MONTH(tr2.date) + YEAR(tr2.date) * 12) <= (month + year * 12)) as cumulative_profit
      FROM rep_transactions tr
               JOIN rep_transaction_groups rtg on tr.group_id = rtg.id
               JOIN rep_transaction_types rtt on rtt.id = tr.type_id
      WHERE rtg.estate_id = ?
        AND tr.date >= ?
      GROUP BY rtg.estate_id, YEAR(tr.date), MONTH(tr.date)
      ORDER BY YEAR(tr.date), MONTH(tr.date)) sub`

	params = append(params, estateId, estateId, estateId, estateId, dateStart)

	rows, err := u.db.Query(query, params...)
	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Println(err.Error())
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
