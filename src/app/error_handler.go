package app

import (
	"encoding/json"
	"errors"
	"net/http"
)

type errorHandler struct{
	err error
	status int
}

var (
	InternalErr = &errorHandler{status: http.StatusInternalServerError, err: errors.New("an error occurred. Try again later")}
)

func (e *errorHandler) ServeHttp(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.status)

	json.NewEncoder(w).Encode(map[string]string{"message": e.err.Error()})
}