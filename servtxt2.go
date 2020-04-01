package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	curr "github.com/vladimirvivien/go-networking/currency/lib0"
)

var currencies = curr.Load("../data.csv")
