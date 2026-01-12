package models

import "time"

type Event struct {
	ID          int    `binding:"required"`
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e *Event) Save() {
	// events = append(events, *e)
	query := `
	INSERT INTO events(id, name, description, location, datetime, userid)
	VALUES(?,?,?,?,?,?)`
	db.DB.prepare(query)
}

func GetAllEvents() []Event {
	return events
}
