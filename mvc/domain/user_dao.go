package domain

import (
	"fmt"
	"net/http"

	"github.com/ricardovegamx/golang-microservices/mvc/utils"
)

/*
Let's assume this is the database.
*/
var users = map[int64]*User{
	123: {Id: 123, FirstName: "Ricardo", LastName: "Vega", Email: "ricardo@ricardovega.pro"},
}

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
