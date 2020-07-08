// Package api provides an example on how to use go-fuzz.
package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Need a named type for our user.
type user struct {
	Type string
	Name string
	Age  int
}

// Routes initializes the routes.
func Routes() {
	http.HandleFunc("/process", Process)
}

// SendError responds with an error.
func SendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(struct{ Error string }{err.Error()})
}

// Process handles the processing of data.
func Process(w http.ResponseWriter, r *http.Request) {

	// Capture the data that was posted over.
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendError(w, err)
		return
	}

	// Split the data by comma.
	parts := strings.Split(string(data), ",")

	// Create a slice of users.
	var users []user

	// Iterate over the set of users we received.
	for _, part := range parts {

		// Extract the user.
		u, err := extractUser(part)
		if err != nil {
			SendError(w, err)
			return
		}
