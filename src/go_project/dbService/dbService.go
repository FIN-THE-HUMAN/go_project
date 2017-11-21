package dbService

import (
	"database/sql"
	"go_project/pavilions"
	"go_project/shelters"

	_ "github.com/lib/pq"
)

func InsertSelter(db *sql.DB, id int, address string) (int64, error) {
	res, err := db.Exec("INSERT INTO shelter VALUES ($1, $2)", id, address)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func InsertPavilions(db *sql.DB, pavilionNumber int, name string, shelterID int, count int) (int64, error) {
	res, err := db.Exec("INSERT INTO pavilions VALUES ($1,$2, $3, $4)",
		pavilionNumber, name, shelterID, count)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SelectAllPavilions(db *sql.DB) ([]pavilions.Pavilion, error) {
	rows, err := db.Query("SELECT * FROM pavilions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]pavilions.Pavilion, 0)
	var pavilion pavilions.Pavilion
	for rows.Next() {
		err = rows.Scan(&pavilion.Number, &pavilion.Name, &pavilion.ID, &pavilion.Count)
		if err != nil {
			return nil, err
		}
		result = append(result, pavilion)
	}
	return result, nil
}

func SelectAllShelters(db *sql.DB) ([]shelters.Shelter, error) {
	rows, err := db.Query("SELECT * FROM shelter")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]shelters.Shelter, 0)
	var shelter shelters.Shelter
	for rows.Next() {
		err = rows.Scan(&shelter.ID, &shelter.Address)
		if err != nil {
			return nil, err
		}
		result = append(result, shelter)
	}
	return result, nil
}
