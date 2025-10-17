package repositories

import (
	"database/sql"
    "fmt"
    "log"
    "encoding/json"

    "github.com/Parovozzzik/real-estate-portfolio/internal/models"
)

type EstateRepository struct {
	db *sql.DB
}

func NewEstateRepository(db *sql.DB) *EstateRepository {
	return &EstateRepository{db: db}
}

func (u *EstateRepository) GetEstates() ([]byte, error) {
    rows, err := u.db.Query(
        "SELECT id, name, estate_type_id, user_id FROM real_estate_portfolio.rep_estates")
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

func (u *EstateRepository) CreateEstate(createEstate *models.CreateEstate) (int64, error) {
    result, err := u.db.Exec(
        "INSERT INTO real_estate_portfolio.rep_estates (name, estate_type_id, user_id, active) VALUES (?, ?, ?, ?)",
        createEstate.Name,
        createEstate.EstateTypeId,
        createEstate.UserId,
        1)
    if err != nil {
        return 0, err
    }

    lastInsertID, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

	return lastInsertID, nil
}