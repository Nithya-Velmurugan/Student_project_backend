package api

import (
	"net/http"
	"student-service/services/app"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := app.InitApp()
	router.ServeHTTP(w, r)
}
