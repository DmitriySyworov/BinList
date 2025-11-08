package storage

import (
	"BinList/app/bins"
	"BinList/app/files"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Storage struct {
	Bins     []bins.Bin `json:"Bins"`
	UpdateAt time.Time  `json:"updateAt"`
}

func NewStorage() (*Storage, string, error) {
	var nameFiles string
	fmt.Println("Укажите название файла, к которому вы хотите добавить бинлист или же создвать новый файл, заполнив с нуля. (Обязательно в формате JSON)")
	fmt.Scan(&nameFiles)
	isMatched := strings.HasSuffix(nameFiles, ".json")
	if !isMatched {
		return nil, "", errors.New("Вы указали неправильный формат. Заполните по следующему образцу: nameFile.json")
	}
	data, err := files.ReadFile(nameFiles)
	if err != nil {
		fmt.Println("Создаем новый файл:", nameFiles)
		return &Storage{
			[]bins.Bin{},
			time.Now(),
		}, nameFiles, nil
	}
	fmt.Println("Добавляем Бин в уже существующий файл:", nameFiles)
	var bins Storage
	err2 := json.Unmarshal(data, &bins)
	if err2 != nil {
		return nil, nameFiles, nil
	}
	return &bins, nameFiles, nil
}
func (bins *Storage) AddStorage(bin bins.Bin, name string) {
	bins.Bins = append(bins.Bins, bin)
	bins.UpdateAt = time.Now()
	file, err := json.Marshal(bins)
	if err != nil {
		fmt.Println(err)
		return
	}
	files.WriteFiles(name, file)
}
func NewReadFile() (*Storage, error) {
	binss, err := ToBytes()
	if err != nil {
		return nil, err
	}
	return binss, nil
}
func ToBytes() (*Storage, error) {
	var nameFiles string
	fmt.Println("Укажите название файла, из которого выхотите прочитать бины. (Обязательно в формате JSON)")
	fmt.Scan(&nameFiles)
	isMatched := strings.HasSuffix(nameFiles, ".json")
	if !isMatched {
		return nil, errors.New("Вы указали неправильный формат. Заполните по следующему образцу: nameFile.json")
	}
	data, err := files.ReadFile(nameFiles)
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
