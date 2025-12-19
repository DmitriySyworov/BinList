package api

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var CaseCreated = []struct {
	nameTest    string
	name        string
	file        string
	status      string
	expectedErr error
}{
	{nameTest: "correct", name: "Dmitriy", file: "first.json", status: "false", expectedErr: nil},
	{nameTest: "incorrectName", name: "", file: "second.json", status: "true", expectedErr: ErrNotName},
	{nameTest: "incorrectFile", name: "Alex", file: "", status: "false", expectedErr: ErrNotFile},
	{nameTest: "incorrectStatus", name: "Dmitriy", file: "third.json", status: "", expectedErr: ErrStatus},
}

func TestCreatedBin(t *testing.T) {
	errEnv := godotenv.Load("/home/dmitriy/GO_BinList/BinList/BinList/.env")
	if errEnv != nil {
		t.Error("Переменные окружения не считаны")
	}
	for _, test := range CaseCreated {
		t.Run(test.nameTest, func(t *testing.T) {
			t.Parallel()
			a := &Api{
				MasterKey: os.Getenv("Master"),
				AccessKey: os.Getenv("Access"),
			}
			ok, err := a.CreatedBin(test.name, test.file, test.status)
			if !ok && err != test.expectedErr {
				t.Errorf("Ожидалось ошибка %v, получаем %v", test.expectedErr, err)
			}
		})
	}
}

var CaseUpdated = []struct {
	name        string
	file        string
	id          string
	status      string
	expectedErr error
}{
	{name: "correct", file: "first.json", id: "694585d8d0ea881f403492bc", status: "true", expectedErr: nil},
	{name: "notFile", file: "", id: "69457456d0ea881f40347683", status: "true", expectedErr: ErrNotFile},
	{name: "notId", file: "first.json", id: "", status: "true", expectedErr: ErrId},
	{name: "notStatus", file: "first.json", id: "69457456d0ea881f40347683", status: "truejsdsjdjs", expectedErr: ErrStatus},
}

func TestUpdateBin(t *testing.T) {
	errEnv := godotenv.Load("/home/dmitriy/GO_BinList/BinList/BinList/.env")
	if errEnv != nil {
		t.Error("Переменные окружения не считаны")
	}
	for _, test := range CaseUpdated {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			a := &Api{
				MasterKey: os.Getenv("Master"),
				AccessKey: os.Getenv("Access"),
			}
			ok, err := a.UpdateBin(test.id, test.file, test.status)
			if !ok && err != test.expectedErr {
				t.Errorf("Ожидалась ошибка %v, получаем %v", test.expectedErr, err)
			}
		})
	}
}

var CaseGetAndDelete = []struct {
	name        string
	id          string
	expectedErr error
}{
	{name: "correct", id: "694585d8d0ea881f403492bc", expectedErr: nil},
	{name: "notId", id: "", expectedErr: ErrId},
	{name: "incorrectId", id: "1", expectedErr: Err400},
}

func TestGetBin(t *testing.T) {
	errEnv := godotenv.Load("/home/dmitriy/GO_BinList/BinList/BinList/.env")
	if errEnv != nil {
		t.Error("Переменные окружения не считаны")
	}
	for _, test := range CaseGetAndDelete {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			a := &Api{
				MasterKey: os.Getenv("Master"),
				AccessKey: os.Getenv("Access"),
			}
			err := a.GetBin(test.id)
			if err != test.expectedErr{
				t.Errorf("Ожидалась ошибка %v, получаем %v", test.expectedErr, err)
			}
		})
	}
}
func TestDeleteBin(t *testing.T){
	errEnv := godotenv.Load("/home/dmitriy/GO_BinList/BinList/BinList/.env")
	if errEnv != nil {
		t.Error("Переменные окружения не считаны")
	}
	for _, test := range CaseGetAndDelete {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			a := &Api{
				MasterKey: os.Getenv("Master"),
				AccessKey: os.Getenv("Access"),
			}
			err := a.DeleteBin(test.id)
			if err != test.expectedErr{
				t.Errorf("Ожидалась ошибка %v, получаем %v", test.expectedErr, err)
			}
		})
	}
}