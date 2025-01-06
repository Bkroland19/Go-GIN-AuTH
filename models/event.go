package models

import (
	"time"

	"example.com/rest-api/db"
)



type Event struct {
ID 			int64
Name 		string `binding:"required"`
Description string `binding:"required"`
Location 	string `binding:"required"`
DateTime 	time.Time `binding:"required"`
UserId 		int
}



var events []Event = []Event{}

func (e Event) Save() error {

	query := `
	INSERT INTO events(name,description,location,dateTime,user_id)
	VALUES (?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name , e.Description,e.Location,e.DateTime,e.UserId)
	if err != nil {
		return err
	}

	id,err := result.LastInsertId()
	e.ID = id
	return nil
}

func GetAllEvents() ([]Event ,error){
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil,err
	}

	defer rows.Close()

	var events []Event = []Event{}

	for rows.Next(){
		var e Event

		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)



		if err != nil {
			return nil,err
		}

		events = append(events, e)
	}
	return events,nil
}



func GetEventById(id int64) (*Event,error){
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query,id)

	var e Event

	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)

	if err != nil {
		return nil,err
	}

	return &e,nil
}