package utils

import (
	"encoding/json"
	"net/http"
)

func HttpErrorHandler(err error, httpStatusCode int, message string, w http.ResponseWriter) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
}
