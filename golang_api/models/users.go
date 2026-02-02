package models

import (
	"errors"
	"fmt"
	"log"
	"udemy-multi-api-golang/db"
	"udemy-multi-api-golang/utils"
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

	// Convert Password to Hashes
	pass, err := utils.HashPassword(u.Password)
	if err != nil {
		fmt.Println("Error Hashing Password.")
		return err
	}

	result, err := stmt.Exec(u.Email, pass)

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

func (u User) ValidateCreds() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var ret_pass string
	err := row.Scan(&u.ID, &ret_pass)

	if err != nil {
		return errors.New("USer Doesn't exist.")
	}

	isValid := utils.CheckPassword(u.Password, ret_pass)
	if !isValid {
		return errors.New("Invalid Password")
	}

	return nil
}
