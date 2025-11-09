package repositories

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(registration *models.Registration) (int64, error) {
	result, err := u.db.Exec(
		"INSERT INTO real_estate_portfolio.rep_users (email, password) VALUES (?, ?)",
		registration.Email,
		string(registration.Password))
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (u *UserRepository) UpdateUser(updateUser *models.UpdateUser) (int64, error) {
	if updateUser.Name == nil && updateUser.Email == nil && updateUser.Phone == nil {
		return updateUser.Id, nil
	}

	params := []any{}
	query := "UPDATE real_estate_portfolio.rep_users SET "
	if updateUser.Name != nil {
		query += "name = ?, "
		params = append(params, updateUser.Name)
	}
	if updateUser.Email != nil {
		query += "email = ?, "
		params = append(params, updateUser.Email)
	}
	if updateUser.Phone != nil {
		query += "phone = ?, "
		params = append(params, updateUser.Phone)
	}
	query = strings.Trim(query, ", ")
	query += " WHERE id = ?"
	params = append(params, updateUser.Id)

	_, err := u.db.Exec(query, params...)
	if err != nil {
		return 0, err
	}

	return updateUser.Id, nil
}

func (u *UserRepository) LoginUser(login *models.Login) (*models.User, error) {
	userData, err := u.db.Query(
		"SELECT id, name, email, phone, password FROM real_estate_portfolio.rep_users WHERE email = ?",
		login.Email)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	userData.Next()
	defer userData.Close()
	err = userData.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Password)
	if err != nil {
		return nil, err
	}

	checkPasswordHash := models.CheckPasswordHash(login.Password, user.Password)
	if checkPasswordHash == false {
		err := errors.New("something went wrong")
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUserById(id int64) (*models.User, error) {
	userData, err := u.db.Query(
		"SELECT id, email, name, phone FROM real_estate_portfolio.rep_users WHERE id = ?",
		id)
	if err != nil {
		return nil, err
	}
	defer userData.Close()

	user := &models.User{}
	userData.Next()
	err = userData.Scan(&user.Id, &user.Email, &user.Name, &user.Phone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUsers() ([]byte, error) {
	rows, err := u.db.Query(
		"SELECT id, email, name, phone FROM real_estate_portfolio.rep_users")
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

func (u *UserRepository) GetUserEstates(userId int64) ([]byte, error) {
	query :=
		"SELECT re.id, re.name, re.description, re.user_id, re.estate_type_id, re.active, FLOOR(RAND() * 101) as recoupment, ret.name as estate_type_name, ret.icon as estate_type_icon " +
			"FROM real_estate_portfolio.rep_estates re " +
			"JOIN real_estate_portfolio.rep_estate_types ret ON ret.id = re.estate_type_id " +
			"WHERE re.user_id = ?"
	rows, err := u.db.Query(query, userId)
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
func (u *UserRepository) GetUserTransactions(userId int64, estateId *int64, filterTransactions *models.FilterTransactions) (*models.PaginatedResponse, error) {
	params := []any{}
	params = append(params, userId)
	query :=
		"SELECT rt.id as transaction_id, rtg.id as transaction_group_id, " +
			"re.id as estate_id, re.name as estate_name, " +
			"ret.id as estate_type_id, ret.name as estate_type_name, " +
			"rtt.id as transaction_type_id, rtt.name as transaction_type_name, rtt.direction as transaction_type_direction, rtt.regularity as transaction_type_regularity, " +
			"rt.sum, rt.date, rt.comment " +
			"FROM real_estate_portfolio.rep_estates re " +
			"JOIN real_estate_portfolio.rep_estate_types ret ON ret.id = re.estate_type_id " +
			"JOIN real_estate_portfolio.rep_transaction_groups rtg ON rtg.estate_id = re.id " +
			"JOIN real_estate_portfolio.rep_transactions rt ON rt.group_id = rtg.id " +
			"JOIN real_estate_portfolio.rep_transaction_types rtt ON rtt.id = rt.type_id " +
			"WHERE re.user_id = ? "
	if estateId != nil || (filterTransactions != nil && filterTransactions.EstateId != nil) {
		query += "AND re.id = ? "
		params = append(params, estateId)
	}
	if filterTransactions != nil {
		if filterTransactions.EstateTypeId != nil {
			query += "AND re.estate_type_id = ? "
			params = append(params, filterTransactions.EstateTypeId)
		}
		if filterTransactions.TransactionTypeId != nil {
			query += "AND rt.type_id = ? "
			params = append(params, filterTransactions.TransactionTypeId)
		}
		if filterTransactions.TransactionGroupId != nil {
			query += "AND rt.group_id = ? "
			params = append(params, filterTransactions.TransactionGroupId)
		}
		if filterTransactions.TransactionTypeDirection != nil {
			query += "AND rtt.direction = ? "
			params = append(params, filterTransactions.TransactionTypeDirection)
		}
		if filterTransactions.TransactionTypeRegularity != nil {
			query += "AND rtt.regularity = ? "
			params = append(params, filterTransactions.TransactionTypeRegularity)
		}
		if filterTransactions.DateStart != nil {
			query += "AND rt.date >= ? "
			params = append(params, filterTransactions.DateStart)
		}
		if filterTransactions.DateEnd != nil {
			query += "AND rt.date <= ? "
			params = append(params, filterTransactions.DateEnd)
		}

		sortBy := "rt.date"
		sortOrder := "ASC"
		if filterTransactions.SortBy != nil {
			switch *filterTransactions.SortBy {
			case "sum":
				sortBy = "rt.sum"
			case "transaction_type_name":
				sortBy = "rtt.name"
			}

			if filterTransactions.SortOrder != nil {
				switch *filterTransactions.SortOrder {
				case "DESC":
					sortOrder = "DESC"
				}
			}
		}
		query += "ORDER BY " + sortBy + " " + sortOrder + " "
	}

	var rowsCount int
	countQuery := "SELECT COUNT(*) FROM (" + query + ") count"
	err := u.db.QueryRow(countQuery, params...).Scan(&rowsCount)
	if err != nil {
		log.Println(err.Error())
	}

	query += "LIMIT ? OFFSET ? "
	var limit int = 10
	var page int64 = 1
	var offset int64 = 0
	if filterTransactions != nil {
		if filterTransactions.Limit != nil {
			limit = *filterTransactions.Limit

			if filterTransactions.Page != nil {
				page = *filterTransactions.Page
				if page < 1 {
					page = 1
				}
				offset = (page - 1) * int64(limit)
			}
		}
	}
	params = append(params, limit)
	params = append(params, offset)

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
				if col == "transaction_type_regularity" || col == "transaction_type_direction" {
					v = val.(int64) != 0
				} else {
					v = val
				}
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	response := &models.PaginatedResponse{
		Data:       tableData,
		TotalItems: int64(rowsCount),
		Page:       page,
		Limit:      limit,
		TotalPages: int64(rowsCount) / int64(limit),
	}

	return response, nil
}

func (u *UserRepository) GetUserEstate(userId, estateId int64) (*models.FullEstate, error) {
	query :=
		"SELECT re.id, re.name, re.description, re.user_id, re.estate_type_id, re.active, ret.name as estate_type_name, ret.icon as estate_type_icon " +
			"FROM real_estate_portfolio.rep_estates re " +
			"JOIN real_estate_portfolio.rep_estate_types ret ON ret.id = re.estate_type_id " +
			"WHERE re.user_id = ? AND re.id = ?"

	estateData, err := u.db.Query(query, userId, estateId)
	if err != nil {
		return nil, err
	}
	defer estateData.Close()

	estate := &models.FullEstate{}
	estateData.Next()
	err = estateData.Scan(&estate.Id, &estate.Name, &estate.Description, &estate.UserId, &estate.EstateTypeId, &estate.Active, &estate.EstateTypeName, &estate.EstateTypeIcon)
	if err != nil {
		return nil, err
	}

	estate.Recoupment = rand.Intn(101)

	return estate, nil
}

func (u *UserRepository) GetUserEstateValues(userId, estateId int64, filterEstateValues *models.FilterEstateValues) (*models.PaginatedResponse, error) {
	params := []any{}
	params = append(params, userId)
	params = append(params, estateId)

	query :=
		"SELECT re.id, rev.date, rev.income, rev.expense, rev.profit, rev.cumulative_income, " +
			"rev.cumulative_expense, rev.cumulative_profit, rev.roi " +
			"FROM real_estate_portfolio.rep_estate_values rev " +
			"JOIN real_estate_portfolio.rep_estates re ON re.id = rev.estate_id " +
			"WHERE re.user_id = ? AND re.id = ? "

	if filterEstateValues != nil {
		if filterEstateValues.DateStart != nil {
			query += "AND rev.date >= ? "
			params = append(params, filterEstateValues.DateStart)
		}
		if filterEstateValues.DateEnd != nil {
			query += "AND rev.date <= ? "
			params = append(params, filterEstateValues.DateEnd)
		}

		sortBy := "rev.date"
		sortOrder := "ASC"
		if filterEstateValues.SortBy != nil && filterEstateValues.SortOrder != nil {
			switch *filterEstateValues.SortOrder {
			case "DESC":
				sortOrder = "DESC"
			}
		}
		query += "ORDER BY " + sortBy + " " + sortOrder + " "
	}

	var rowsCount int
	countQuery := "SELECT COUNT(*) FROM (" + query + ") count"
	err := u.db.QueryRow(countQuery, params...).Scan(&rowsCount)
	if err != nil {
		log.Println(err.Error())
	}

	query += "LIMIT ? OFFSET ? "
	var limit int = 120
	var page int64 = 1
	var offset int64 = 0
	if filterEstateValues != nil {
		if filterEstateValues.Limit != nil {
			limit = *filterEstateValues.Limit
		} else if filterEstateValues.DateStart != nil && filterEstateValues.DateEnd != nil {
			dateStart, err := time.Parse("2006-01-02", *filterEstateValues.DateStart)
			if err != nil {
				log.Println(err.Error())
			}

			dateEnd, err := time.Parse("2006-01-02", *filterEstateValues.DateEnd)
			if err != nil {
				log.Println(err.Error())
			}

			limit = diffMonths(dateEnd, dateStart)
		}

		if filterEstateValues.Page != nil {
			page = *filterEstateValues.Page
			offset = page * int64(limit)
			if page < 1 {
				page = 1
			}
			offset = (page - 1) * int64(limit)
		}
	}
	params = append(params, limit)
	params = append(params, offset)

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

	response := &models.PaginatedResponse{
		Data:       tableData,
		TotalItems: int64(rowsCount),
		Page:       page,
		Limit:      len(tableData),
		TotalPages: int64(rowsCount) / int64(limit),
	}

	return response, nil
}

func diffMonths(t1, t2 time.Time) int {
	if t1.After(t2) {
		t1, t2 = t2, t1
	}

	year1, month1, _ := t1.Date()
	year2, month2, _ := t2.Date()

	months := 0
	months += int(month2) - int(month1)
	months += (year2 - year1) * 12

	return months
}
