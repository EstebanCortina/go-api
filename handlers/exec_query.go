package handlers

import (
	"crud/exceptions"
	"crud/models"
	"database/sql"
)

func ExecQuery(
	db *sql.DB,
	query string,
	results chan models.User) {
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		panic((exceptions.ServerException{
			Name: "mysql",
			Body: "Error executing query",
			Err:  err,
		}))
	}
	defer rows.Close()

	for rows.Next() {
		var result models.User
		if err := rows.Scan(
			&result.Id,
			&result.Name,
			&result.User_name,
			&result.Followers); err != nil {
			panic((exceptions.ServerException{
				Name: "mysql",
				Body: "Error scanning row",
				Err:  err,
			}))
		}
		results <- result
	}
	close(results)

	if err := rows.Err(); err != nil {
		panic("Error al iterar sobre los resultados")
	}

}
