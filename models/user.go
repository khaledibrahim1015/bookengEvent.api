package models

import (
	"errors"

	"bookingEvent.api/db"
	"bookingEvent.api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {

	query := `
	INSERT INTO users (email , password )
	VALUES (? ,  ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userid, err := result.LastInsertId()
	u.ID = userid
	return err

}

func (usr *User) ValidateCredentails() error {

	//  select current user
	query := "SELECT id,password FROM users where email =?"

	row := db.DB.QueryRow(query, usr.Email)

	var retreviedHashedPassword string
	err := row.Scan(&usr.ID, &retreviedHashedPassword)
	if err != nil {
		return err
	}

	// validate password
	passwordIsValid := utils.CheckPasswordHash(usr.Password, retreviedHashedPassword)
	if !passwordIsValid {
		return errors.New("credentails not Vaild ")
	}
	return nil

}
