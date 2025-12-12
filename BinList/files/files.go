package files

import (
	"errors"
	"os"
	"strings"
)

func Read(fileName string) ([]byte, error) {
	isMatched := strings.HasSuffix(fileName, ".json")
	if !isMatched {
		return nil, errors.New("Вы указали неправильный формат. Заполните по следующему образцу: nameFile.json")
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, nil
	}
	return data,  nil
}
func Write(fileName string, data []byte) error{
	isMatched := strings.HasSuffix(fileName, ".json")
	if !isMatched {
		return errors.New("Вы указали неправильный формат. Заполните по следующему образцу: nameFile.json")
	}
	files, err1 := os.Create(fileName)
	if err1 != nil {
		return err1
	}
	defer files.Close()
	_, err2 := files.Write(data)
	if err2 != nil {
		return err2
	}
	return nil
}
