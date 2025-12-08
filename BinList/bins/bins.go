package bins

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	Name      string    `json:"name"`
	Id        string    `json:"ID"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewBinCreate(name string) (*Bin, error) {
	for _, value := range name {
		if (value < 'a' || value > 'z') && (value < 'A' || value > 'Z') {
			return nil, errors.New("Указано не корректное имя")
		}
	}
	return &Bin{
		Name:      name,
		Id:        "",
		CreatedAt: time.Now(),
	}, nil
}
func NewBinUpdate(id string) (*Bin, error) {
	return &Bin{
		Name:      "",
		Id:        id,
		CreatedAt: time.Now(),
	}, nil
}
func (bin *Bin) OutputBins() {
	fmt.Printf("name: %s\n ID: %s\n createdAt: %s\n", bin.Name, bin.Id, bin.CreatedAt)
}
