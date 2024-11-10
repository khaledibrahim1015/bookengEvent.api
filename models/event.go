package models

import (
	"time"

	"bookingEvent.api/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      int
}

func (ev *Event) Save() error {

	query := `
	INSERT INTO events (name , description , location , datetime , user_id)
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(ev.Name, ev.Description, ev.Location, ev.DateTime, ev.UserId)
	if err != nil {
		return err
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	ev.ID = insertedId
	return nil

}

func GetAllEvents() ([]Event, error) {

	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//  map
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {

	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		// return Event{}, err
		return nil, err
	}
	return &event, nil
}
