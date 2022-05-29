package utils

import (
	"encoding/json"
	"net/http"
)

/* HTTP Responders */
type response struct{}

var Respond response

/*
This function responds the information successfully as JSON.
*/
func (r response) JSON(w http.ResponseWriter, code int, payload interface{}) {
	res, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}

/*
This function responds an error message to the frontend as JSON.
*/
func (r response) Error(w http.ResponseWriter, code int, message string) {
	r.JSON(w, code, map[string]string{"error": message})
}
