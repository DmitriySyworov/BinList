package bins

import (
	"math/rand"
	"time"
)

type Bin struct {
	Name      string    `json:"Name"`
	Id        string    `json:"ID"`
	Password  string    `json:"password"`
	Private   string    `json:"private"`
	LocalFile string    `json:"localFile"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewBin(name, id, password, private, file string) (*Bin, error) {
	newPassword := ""
	if password == "" {
		newPassword = makePassword()
	} else {
		newPassword = password
	}

	return &Bin{
		Name:      name,
		Id:        id,
		Password:  newPassword,
		Private:   private,
		LocalFile: file,
		CreatedAt: time.Now(),
	}, nil
}
func makePassword() string {
	var trueRand []byte
	stop := 0
	for {
		random := rand.Intn(127)
		if random > 32 {
			trueRand = append(trueRand, byte(random))
			stop++
		}
		if stop == 20 {
			break
		}
	}
	return string(trueRand)
}
