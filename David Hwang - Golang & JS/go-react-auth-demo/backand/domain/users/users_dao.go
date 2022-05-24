package users

import (
	"go-react-auth-demo/datasource/mysql/users_db"
	"go-react-auth-demo/utils/errors"
)

var (
	queryInsertUser     = "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)"
	queryGetUserByEmail = "SELECT id, first_name, last_name, email, password FROM users WHERE email=?"
	queryGetUserByID    = "SELECT id, first_name, last_name, email FROM users WHERE id=? "
)

func (u *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewBadRequestError("Database Error")
	}

	defer stmt.Close()

	res, err := stmt.Exec(u.Firstname, u.Lastname, u.Email, u.Password)

	if err != nil {
		return errors.NewInternalServerError("Database Error")
	}

	userID, err := res.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError("Database Error")
	}

	u.ID = userID

	return nil

}

func (user *User) GetByEmail() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUserByEmail)

	if err != nil {
		return errors.NewInternalServerError("Invalid Email")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Email)

	if getErr := result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Password); getErr != nil {
		return errors.NewInternalServerError("Database Error")
	}

	return nil
}

func (u *User) GetByID() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUserByID)

	if err != nil {
		return errors.NewInternalServerError("Database Error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.ID)

	if getErr := result.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email); getErr != nil {
		return errors.NewInternalServerError("Database Error")
	}

	return nil

}
