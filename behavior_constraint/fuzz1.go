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
