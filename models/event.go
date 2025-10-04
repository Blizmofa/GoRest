package models

import (
	"time"

	"example.com/go-rest/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

const (
	sqlInsertEvent = `
		INSERT INTO events(name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)`

	sqlSelectAllEvents = `
		SELECT * FROM events`

	sqlSelectEventByID = `
		SELECT * FROM events WHERE id = ?`

	sqlUpdateEventByID = `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?`

	sqlDeleteEventByID = `
		DELETE FROM events WHERE id = ?`

	sqlInsertRegistration = `
		INSERT INTO registrations(event_id, user_id) VALUES (?, ?)`

	sqlDeleteRegistration = `
		DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
)

func (event *Event) Save() error {

	result, err := db.ExecStatement(sqlInsertEvent, event.Name, event.Description, event.Location, event.DateTime, event.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {

	rows, err := db.DB.Query(sqlSelectAllEvents)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {

	row := db.DB.QueryRow(sqlSelectEventByID, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event *Event) Update() error {
	_, err := db.ExecStatement(sqlUpdateEventByID, event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

func (event *Event) Delete() error {

	_, err := db.ExecStatement(sqlDeleteEventByID, event.ID)
	return err
}

func (event *Event) Register(userId int64) error {

	_, err := db.ExecStatement(sqlInsertRegistration, event.ID, userId)
	return err
}

func (event *Event) CancelRegistration(userId int64) error {

	_, err := db.ExecStatement(sqlDeleteRegistration, event.ID, userId)
	return err
}
