package main

import (
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	createdAt time.Time
	name      string
}

var BinList = []Bin{}

func newBin() (*Bin, error) {
}
func main() {

}
