package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	curr "github.com/vladimirvivien/go-networking/currency/lib"
)

var (
	currencies = curr.Load("../data.csv")
)
