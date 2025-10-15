package repositories

import (
	"database/sql"
	"errors"
    "fmt"
    "log"
    "encoding/json"

    "github.com/Parovozzzik/real-estate-portfolio/internal/logging"
    "github.com/Parovozzzik/real-estate-portfolio/internal/models"
)

/* type IUserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetUsers(id string) (*models.User, error)
	Auth(userName string, password string) (*models.User, error)
} */

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(registration *models.Registration) (int64, error) {
    result, err := u.db.Exec("INSERT INTO real_estate_portfolio.users (email, password) VALUES (?, ?)", registration.Email, string(registration.Password))
    if err != nil {
        return 0, err
    }

    lastInsertID, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

	return lastInsertID, nil
}

func (u *UserRepository) LoginUser(login *models.Login) (*models.User, error) {
    hashedPassword, err := models.HashPassword(login.Password)

        logging.Init()
    logger := logging.GetLogger()
                logger.Println(hashedPassword)


    if err != nil {
        return nil, err
    }

    userData, err := u.db.Query("SELECT id, email, password FROM real_estate_portfolio.users WHERE email = ?", login.Email)
    if err != nil {
        return nil, err
    }

    user := &models.User{}
    userData.Next()
    err = userData.Scan(&user.Id, &user.Email, &user.Password)
    if err != nil {
        return nil, err
    }

    checkPasswordHash := models.CheckPasswordHash(login.Password, user.Password)
    if (checkPasswordHash == false) {
        err := errors.New("something went wrong")
        return nil, err
    }

	return user, nil
}

func(u *UserRepository) GetUserByID(id string) (*models.User, error) {
	user := models.NewUser(1, "username", "email@email.ru", "password")
    return user, nil
}

func (u *UserRepository) GetUsers() ([]byte, error) {
    rows, err := u.db.Query("SELECT id, email, name FROM real_estate_portfolio.users")
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