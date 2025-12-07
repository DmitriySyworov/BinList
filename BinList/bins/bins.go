package bins

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	Name      string    `json:"name"`
	Id        string    `json:"ID"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewBin(name, id string, private bool) (*Bin, error) {
	for _, value := range name {
		if (value < 'a' || value > 'z') && (value < 'A' || value > 'Z') {
			return nil, errors.New("Указано не корректное имя")
		}
	}
	if len(id) != 6 {
		return nil, errors.New("ID указан некорректно: ID должен состоять из 6 символов")
	}
	return &Bin{
		Name:      name,
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
	}, nil
}

func (bin *Bin) OutputBins() {
	fmt.Printf("name: %s\n ID: %s\n private: %t\n createdAt: %s\n", bin.Name, bin.Id, bin.Private, bin.CreatedAt)
}
