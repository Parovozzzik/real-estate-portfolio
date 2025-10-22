package repositories

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Parovozzzik/real-estate-portfolio/internal/logging"
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

	logging.Init()
	logger := logging.GetLogger()
	logger.Println(query)
	logger.Println(params)

	_, err := u.db.Exec(query, params...)
	if err != nil {
		return 0, err
	}

	return updateUser.Id, nil
}

func (u *UserRepository) LoginUser(login *models.Login) (*models.User, error) {
	userData, err := u.db.Query(
		"SELECT id, email, password FROM real_estate_portfolio.rep_users WHERE email = ?",
		login.Email)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	userData.Next()
	defer userData.Close()
	err = userData.Scan(&user.Id, &user.Email, &user.Password)
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
		"SELECT id, email, name FROM real_estate_portfolio.rep_users WHERE id = ?",
		id)
	if err != nil {
		return nil, err
	}
	defer userData.Close()

	user := &models.User{}
	userData.Next()
	err = userData.Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUsers() ([]byte, error) {
	rows, err := u.db.Query(
		"SELECT id, email, name FROM real_estate_portfolio.rep_users")
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
		"SELECT re.id, re.name, re.user_id, re.estate_type_id, re.active, ret.name as estate_type_name, ret.icon as estate_type_icon " +
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

func (u *UserRepository) GetUserEstate(userId, estateId int64) (*models.FullEstate, error) {
	query :=
		"SELECT re.id, re.name, re.user_id, re.estate_type_id, re.active, ret.name as estate_type_name, ret.icon as estate_type_icon " +
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
	err = estateData.Scan(&estate.Id, &estate.Name, &estate.UserId, &estate.EstateTypeId, &estate.Active, &estate.EstateTypeName, &estate.EstateTypeIcon)
	if err != nil {
		return nil, err
	}

	return estate, nil
}
