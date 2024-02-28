package utils

import (
	errorhandling "articleproject/error"
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorGenerator(w http.ResponseWriter, err error) {
	var response interface{}

	fmt.Println(err.Error())

	if error, ok := err.(errorhandling.CustomError); ok {
		response = errorhandling.CustomError{
			StatusCode: error.StatusCode,
			Message:    error.Message,
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(error.StatusCode)

	} else {
		response = errorhandling.CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(response)
}
