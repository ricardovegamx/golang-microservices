package app

import (
	"net/http"

	"github.com/ricardovegamx/golang-microservices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		panic(err)
	}
}
