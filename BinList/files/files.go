package files

import (
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	isMatched := strings.HasSuffix(name, ".json")
	if !isMatched {
		panic("Вы указали неправильный формат. Заполните по следующему образцу: nameFile.json")
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func WriteFiles(name string, data []byte) error {
	isMatched := strings.HasSuffix(name, ".json")
	if !isMatched {
		panic("Вы указали неправильный формат. Заполните по следующему образцу: nameFile.json")
	}
	files, err1 := os.Create(name)
	if err1 != nil {
		return err1
	}
	_, err2 := files.Write(data)
	if err2 != nil {
		return err2
	}
	return nil
}
