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
