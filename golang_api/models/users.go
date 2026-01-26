package models

import (
	"fmt"
	"log"
	"udemy-multi-api-golang/db"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password)
		VALUES(?,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		log.Println("Error Writing the Data to DB")
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return nil
}

func (u User) Delete() (int64, error) {
	query := `DELETE from users where email = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error in the delete statement.")
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.ID)
	if err != nil {
		fmt.Println("Error Deleting event.")
		return 0, err
	}

	return result.RowsAffected()
}
