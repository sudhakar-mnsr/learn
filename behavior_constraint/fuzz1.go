// Package api provides an example on how to use go-fuzz.
package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)
