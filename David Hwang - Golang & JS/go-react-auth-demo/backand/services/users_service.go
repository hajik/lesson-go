package services

import (
	"go-react-auth-demo/domain/users"

	api_err "go-react-auth-demo/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *api_err.RestError) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	// *encrypt the Password
	passSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return nil, api_err.NewBadRequestError("Failed to encrypt password")
	}

	user.Password = string(passSlice[:])

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil

}

func GetUser(user users.User) (*users.User, *api_err.RestError) {

	result := &users.User{Email: user.Email}

	if err := result.GetByEmail(); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return nil, api_err.NewBadRequestError("Failed to decrypt the password")
	}

	resultWp := &users.User{ID: result.ID, Firstname: result.Firstname, Lastname: result.Lastname, Email: result.Email}

	return resultWp, nil

}

func GetUserByID(userId int64) (*users.User, *api_err.RestError) {
	result := &users.User{ID: userId}

	if err := result.GetByID(); err != nil {
		return nil, err
	}

	return result, nil
}
