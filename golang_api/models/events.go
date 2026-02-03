package models

import (
	"fmt"
	"log"
	"time"
	"udemy-multi-api-golang/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int64     `json:"userId"`
}

func (e *Event) Save() error {
	// events = append(events, e)
	query := `INSERT INTO events(name, description, location, dateTime, user_id)
		VALUES(?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location,
		e.DateTime, e.UserID)

	if err != nil {
		log.Println("Error Writing the Data to DB")
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return nil
}

func GetAll() ([]Event, error) {
	var events []Event
	query := `Select * from events`
	res, err := db.DB.Query(query)
	if err != nil {
		fmt.Println("Error Getting Data from the Database.", err)
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var event Event
		err := res.Scan(&event.ID, &event.Name, &event.Description,
			&event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			fmt.Println("Error reading data")
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil

}
func GetId(id int64) (*Event, error) {
	// We user pointer to an event as the default value of a Pointer can be a nil.
	query := `SELECT * from events where id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description,
		&event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		fmt.Println("Error reading data")
		return nil, err
	}
	return &event, nil
}

func (e Event) Update() error {
	query := `UPDATE EVENTS	SET name = ?, 
	description = ?, location = ?, dateTime = ?
	where id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error Preparing the Statement.")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		fmt.Println("Error Updating values")
		return err
	}
	return nil
}

func (e Event) Delete() (int64, error) {
	query := `DELETE from events where id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error in the delete statement.")
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.ID)
	if err != nil {
		fmt.Println("Error Deleting event.")
		return 0, err
	}

	return result.RowsAffected()
}
