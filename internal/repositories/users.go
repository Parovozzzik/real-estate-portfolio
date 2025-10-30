package repositories

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"

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

func (u *UserRepository) GetUserTransactions(userId int64) ([]byte, error) {
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
			"WHERE re.user_id = ? " +
			"ORDER BY rt.date"
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
