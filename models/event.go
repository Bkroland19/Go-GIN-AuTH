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
UserId 		int64
}



var events []Event = []Event{}

// Save inserts a new event into the events table in the database.
// It prepares an SQL INSERT statement and executes it with the event's attributes.
// If the insertion is successful, it updates the Event struct with the new event's ID.
func (e Event) Save() error {
    // Define the SQL INSERT query to add a new event to the events table.
    query := `
    INSERT INTO events(name,description,location,dateTime,user_id)
    VALUES (?,?,?,?,?)
    `

    // Prepare the SQL statement for execution.
    stmt, err := db.DB.Prepare(query)
    if err != nil {
        // If there is an error preparing the statement, return the error.
        return err
    }
    // Ensure the statement is closed after the function completes to free up resources.
    defer stmt.Close()

    // Execute the prepared statement with the event's attributes.
    result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
    if err != nil {
        // If there is an error executing the statement, return the error.
        return err
    }

    // Retrieve the ID of the newly inserted event.
    id, err := result.LastInsertId()
    if err != nil {
        // If there is an error retrieving the last insert ID, return the error.
        return err
    }

    // Update the Event struct with the new event's ID.
    e.ID = id

    // Return nil to indicate the operation was successful.
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

func (e Event) Update() error{
	query := `UPDATE events SET name=?, description=?, location=?, dateTime=? WHERE id=?`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	if err != nil {
		return err
	}

	return nil
}

func (e Event) Delete() error{
	query := `DELETE FROM events WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
		}
	return nil
}


func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id,user_id) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
     return err
	}

	defer stmt.Close()


	_, err = stmt.Exec(e.ID,userId)

	return err
}