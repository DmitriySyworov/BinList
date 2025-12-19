package api

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	dir, _ := os.Getwd()
	file := filepath.Join(dir, ".env")
	errEnv := godotenv.Load(file)
	if errEnv != nil {
		panic("Переменные окружения не считаны")
	}
	code := m.Run()
	os.Exit(code)
}

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
	for _, test := range CaseCreated {
		t.Run(test.nameTest, func(t *testing.T) {
			t.Parallel()
			a := &Api{
				MasterKey: os.Getenv("Master"),
				AccessKey: os.Getenv("Access"),
			}
			data, err := a.CreatedBin(test.name, test.file, test.status, "")
			if err != test.expectedErr {
				t.Errorf("Ожидалось ошибка %v, получаем %v", test.expectedErr, err)
			}
			var creatResp CreatedResponce
			json.Unmarshal(data, &creatResp)
			defer a.DeleteBin(creatResp.Metadata.Id)
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
	{name: "notFile", file: "", id: "69457456d0ea881f40347683", status: "true", expectedErr: ErrNotFile},
	{name: "notId", file: "first.json", id: "", status: "true", expectedErr: ErrId},
	{name: "notStatus", file: "first.json", id: "69457456d0ea881f40347683", status: "truejsdsjdjs", expectedErr: ErrStatus},
}

func TestUpdateBinNegative(t *testing.T) {
	for _, test := range CaseUpdated {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			a := &Api{
				MasterKey: os.Getenv("Master"),
				AccessKey: os.Getenv("Access"),
			}
			err := a.UpdateBin(test.id, test.file, test.status, "")
			if err != test.expectedErr {
				t.Errorf("Ожидалась ошибка %v, получаем %v", test.expectedErr, err)
			}

		})
	}
}
func TestUpdateBin(t *testing.T) {
	a, creatResp, errCreat := createrTester()
	if errCreat != nil {
		t.Errorf("Ожидалось успешное создание файла, получаем ошибку %v", errCreat)
	}
	errTrue := a.UpdateBin(creatResp.Metadata.Id, "first.json", "true", "")
	if errTrue != nil {
		t.Errorf("Ожидалась удачное обновление файла, но мы получаем: %v", errTrue)
	}

	defer a.DeleteBin(creatResp.Metadata.Id)
}

var CaseGetAndDelete = []struct {
	name        string
	id          string
	expectedErr error
}{
	{name: "notId", id: "", expectedErr: ErrId},
	{name: "incorrectId", id: "1", expectedErr: Err400},
}

func TestGetBinNegative(t *testing.T) {
	for _, test := range CaseGetAndDelete {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			a := &Api{
				MasterKey: os.Getenv("Master"),
				AccessKey: os.Getenv("Access"),
			}
			err := a.GetBin(test.id)
			if err != test.expectedErr {
				t.Errorf("Ожидалась ошибка %v, получаем %v", test.expectedErr, err)
			}
		})
	}
}
func TestGetBin(t *testing.T) {
	a, creatResp, errCreat := createrTester()
	if errCreat != nil {
		t.Errorf("Ожидалось успешное создание файла, получаем ошибку %v", errCreat)
	}
	err := a.GetBin(creatResp.Metadata.Id)
	if err != nil {
		t.Errorf("Ожидалось удачное выполнение функции Get, получаем: %v", err)
	}
	defer a.DeleteBin(creatResp.Metadata.Id)
}
func TestDeleteBinNegative(t *testing.T) {
	for _, test := range CaseGetAndDelete {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			a := &Api{
				MasterKey: os.Getenv("Master"),
				AccessKey: os.Getenv("Access"),
			}
			err := a.DeleteBin(test.id)
			if err != test.expectedErr {
				t.Errorf("Ожидалась ошибка %v, получаем %v", test.expectedErr, err)
			}
		})
	}
}
func TestDeleteBin(t *testing.T) {
	a, creatResp, errCreat := createrTester()
	if errCreat != nil {
		t.Errorf("Ожидалось успешное создание файла, получаем ошибку %v", errCreat)
	}
	err := a.DeleteBin(creatResp.Metadata.Id)
	if err != nil {
		t.Errorf("Ожидалось удачное удаление, получаем ошибку: %v", err)
	}
}
func createrTester() (*Api, *CreatedResponce, error) {
	a := &Api{
		MasterKey: os.Getenv("Master"),
		AccessKey: os.Getenv("Access"),
	}
	data, errCreat := a.CreatedBin("Dmitriy", "first.json", "true", "")
	if errCreat != nil {
		return nil, nil, errCreat
	}
	var creatResp CreatedResponce
	json.Unmarshal(data, &creatResp)
	return a, &creatResp, nil
}
