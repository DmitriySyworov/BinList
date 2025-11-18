package files

import (
	"errors"
	"os"
	"strings"
)

type JsonDb struct {
	FileName string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		FileName: name,
	}
}
func (db *JsonDb) Read() ([]byte, string, error, error) {
	isMatched := strings.HasSuffix(db.FileName, ".json")
	if !isMatched {
		return nil, db.FileName, nil, errors.New("Вы указали неправильный формат. Заполните по следующему образцу: nameFile.json")
	}
	data, err := os.ReadFile(db.FileName)
	if err != nil {
		return nil, db.FileName, err, nil
	}
	return data, db.FileName, nil, nil
}
func (db *JsonDb) Write(data []byte) (error, error) {
	isMatched := strings.HasSuffix(db.FileName, ".json")
	if !isMatched {
		return nil, errors.New("Вы указали неправильный формат. Заполните по следующему образцу: nameFile.json")
	}
	files, err1 := os.Create(db.FileName)
	if err1 != nil {
		return err1, nil
	}
	_, err2 := files.Write(data)
	if err2 != nil {
		return err2, nil
	}
	return nil, nil
}
