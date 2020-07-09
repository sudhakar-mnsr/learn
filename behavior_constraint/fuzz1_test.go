// Tests to validate the api endpoints.
package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/testing/fuzzing/example1"
)

const succeed = "\u2713"
const failed = "\u2717"

func init() {
	api.Routes()
}

// TestProcess tests the Process endpoint with proper data.
func TestProcess(t *testing.T) {
	tests := []struct {
		url    string
		status int
		val    []byte
		resp   string
	}{
