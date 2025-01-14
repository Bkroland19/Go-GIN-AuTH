package models

import (
	"database/sql"
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID 			int64
	Email       string `binding:"required"`
	Password    string `binding:"required"`
}


func (u *User) Save() error {
    query := `
    INSERT INTO users(email,password) VALUES(?,?)`

    stmt, err := db.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

	hashedPassword,err := utils.HashPassword(u.Password)

	 if err != nil {
        return err
    }

    result, err := stmt.Exec(u.Email, hashedPassword)
    if err != nil {
        return err
    }

    userId, err := result.LastInsertId()
    if err != nil {
        return err
    }

    u.ID = userId

    return nil
}


// ValidateCredentials checks if the provided password matches the stored hashed password for the user.
// It queries the database for the user's email and retrieves the associated hashed password.
// If the password is valid, it returns nil; otherwise, it returns an error indicating invalid credentials.
func (u *User) ValidateCredentials() error {
    query := "SELECT password,id FROM users WHERE email = ?"
    row := db.DB.QueryRow(query, u.Email)

    var retrievedPassword string
    err := row.Scan(&u.ID ,&retrievedPassword)
    if err == sql.ErrNoRows {
        return errors.New("email not found")
    } else if err != nil {
        return err 
    }

    if !utils.CheckPasswordHash(u.Password, retrievedPassword) {
        return errors.New("passwords do not match")
    }

    return nil
}
