package curlib

import (
   "encoding/csv"
   "io"
   "os"
   "strings"
)

type Currency struct {
   Code string
   Name string
   Number string
   Country string
}

func Load(path string) []Currency {
table := make([]Currency, 0)
file, err := os.Open(path)
if err != nil {
   panic(err.Error())
}
defer file.Close()
   
