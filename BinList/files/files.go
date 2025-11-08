package files

import (
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func WriteFiles(name string, data []byte) error {
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
