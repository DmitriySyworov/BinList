package bins

import (
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

var BinList = []Bin{}

func NewBin() (*Bin, error) {}
