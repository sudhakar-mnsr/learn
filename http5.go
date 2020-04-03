package main

import (
	"fmt"
	"net/http"
)

func main() {
	hello := func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Content-Type", "text/html")
		resp.WriteHeader(http.StatusOK)
		fmt.Fprint(resp, "Hello from Above!")
	}
