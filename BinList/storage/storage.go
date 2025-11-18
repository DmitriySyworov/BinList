package storage

import (
	"BinList/app/bins"
	"encoding/json"
	"fmt"
	"time"
)

type Db interface {
	Read() ([]byte, string, error, error)
	Write(data []byte) (error, error)
}
type Storage struct {
	Bins     []bins.Bin `json:"Bins"`
	UpdateAt time.Time  `json:"updateAt"`
}
type StorageWithBd struct {
	Storage
	Db
}

func NewStorage(db Db) (*StorageWithBd, string, error) {
	data, names, err, errorImportant := db.Read()
	if errorImportant != nil {
		return nil, names, errorImportant
	}
	if err != nil {
		fmt.Println("Создаем новый файл:", names)
		return &StorageWithBd{
			Storage: Storage{
				[]bins.Bin{},
				time.Now(),
			},
			Db: db,
		}, names, nil
	}
	fmt.Println("Добавляем Бин в уже существующий файл:", names)
	var bins Storage
	err2 := json.Unmarshal(data, &bins)
	if err2 != nil {
		return nil, names, nil
	}
	return &StorageWithBd{
		Storage: bins,
		Db:      db,
	}, names, nil
}
func (bins *StorageWithBd) AddStorage(bin bins.Bin, name string) {
	bins.Bins = append(bins.Bins, bin)
	bins.UpdateAt = time.Now()
	file, err := json.Marshal(bins)
	if err != nil {
		fmt.Println(err)
		return
	}
	bins.Db.Write(file)
}
func NewReadFile(db Db) (*Storage, error) {
	data, _, err, errorImportant := db.Read()
	if errorImportant != nil {
		return nil, errorImportant
	}
	if err != nil {
		return nil, err
	}
	var binss Storage
	err2 := json.Unmarshal(data, &binss)
	if err2 != nil {
		return nil, err2
	}
	return &binss, nil
}
