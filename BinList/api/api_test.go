package api

import (
	"BinList/app/storage"
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
	{nameTest: "correct", name: "Xzinxzzkes", file: "news.json", status: "false", expectedErr: nil},
	{nameTest: "incorrectName", name: "", file: "seconhhhd.json", status: "true", expectedErr: ErrNotName},
	{nameTest: "incorrectFile", name: "Alexitre", file: "", status: "false", expectedErr: ErrNotFile},
	{nameTest: "incorrectStatus", name: "Dmitvcriy", file: "thirdlmn.json", status: "", expectedErr: ErrStatus},
}

func TestCreatedBin(t *testing.T) {
	for _, test := range CaseCreated {
		t.Run(test.nameTest, func(t *testing.T) {
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
			defer storage.DeletedLocal(creatResp.Metadata.Id)
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
	a := &Api{
		MasterKey: os.Getenv("Master"),
		AccessKey: os.Getenv("Access"),
	}
	data, errCreat := a.CreatedBin("ncxmc", "pllk.json", "true", "")
	if errCreat != nil {
		t.Error(errCreat)
	}
	var creatResp CreatedResponce
	json.Unmarshal(data, &creatResp)
	defer a.DeleteBin(creatResp.Metadata.Id)
	defer storage.DeletedLocal(creatResp.Metadata.Id)
	for _, test := range CaseUpdated {
		t.Run(test.name, func(t *testing.T) {
			err := a.UpdateBin(test.id, test.file, test.status, "", "Alexxc")
			if err != test.expectedErr {
				t.Errorf("Ожидалась ошибка %v, получаем %v", test.expectedErr, err)
			}

		})
	}
}
func TestUpdateBin(t *testing.T) {
	a := &Api{
		MasterKey: os.Getenv("Master"),
		AccessKey: os.Getenv("Access"),
	}
	data, errCreat := a.CreatedBin("jhhgg", "uuiu.json", "true", "")
	if errCreat != nil {
		t.Error(errCreat)
	}
	var creatResp CreatedResponce
	json.Unmarshal(data, &creatResp)
	defer a.DeleteBin(creatResp.Metadata.Id)
	defer storage.DeletedLocal(creatResp.Metadata.Id)
	errTrue := a.UpdateBin(creatResp.Metadata.Id, "wer.json", "true", "", "okli")
	if errTrue != nil {
		t.Errorf("Ожидалась удачное обновление файла, но мы получаем: %v", errTrue)
	}
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
	a := &Api{
		MasterKey: os.Getenv("Master"),
		AccessKey: os.Getenv("Access"),
	}
	data, errCreat := a.CreatedBin("Dhjh", "poii.json", "true", "")
	if errCreat != nil {
		t.Error(errCreat)
	}
	var creatResp CreatedResponce
	json.Unmarshal(data, &creatResp)
	defer a.DeleteBin(creatResp.Metadata.Id)
	defer storage.DeletedLocal(creatResp.Metadata.Id)
	for _, test := range CaseGetAndDelete {
		t.Run(test.name, func(t *testing.T) {
			err := a.GetBin(test.id)
			if err != test.expectedErr {
				t.Errorf("Ожидалась ошибка %v, получаем %v", test.expectedErr, err)
			}
		})
	}
}
func TestGetBin(t *testing.T) {
	a := &Api{
		MasterKey: os.Getenv("Master"),
		AccessKey: os.Getenv("Access"),
	}
	data, errCreat := a.CreatedBin("bbbb", "ccc.json", "true", "")
	if errCreat != nil {
		t.Error(errCreat)
	}
	var creatResp CreatedResponce
	json.Unmarshal(data, &creatResp)
	defer a.DeleteBin(creatResp.Metadata.Id)
	defer storage.DeletedLocal(creatResp.Metadata.Id)
	err := a.GetBin(creatResp.Metadata.Id)
	if err != nil {
		t.Errorf("Ожидалось удачное выполнение функции Get, получаем: %v", err)
	}
}
func TestDeleteBinNegative(t *testing.T) {
	a := &Api{
		MasterKey: os.Getenv("Master"),
		AccessKey: os.Getenv("Access"),
	}
	data, errCreat := a.CreatedBin("Dmitriy2ss", "a.json", "true", "")
	if errCreat != nil {
		t.Error(errCreat)
	}
	var creatResp CreatedResponce
	json.Unmarshal(data, &creatResp)
	defer a.DeleteBin(creatResp.Metadata.Id)
	defer storage.DeletedLocal(creatResp.Metadata.Id)
	for _, test := range CaseGetAndDelete {
		t.Run(test.name, func(t *testing.T) {
			err := a.DeleteBin(test.id)
			if err != test.expectedErr {
				t.Errorf("Ожидалась ошибка %v, получаем %v", test.expectedErr, err)
			}
		})
	}
}
func TestDeleteBin(t *testing.T) {
	a := &Api{
		MasterKey: os.Getenv("Master"),
		AccessKey: os.Getenv("Access"),
	}
	data, errCreat := a.CreatedBin("Dmitrqqiy", "firstyt.json", "true", "")
	if errCreat != nil {
		t.Error(errCreat)
	}
	var creatResp CreatedResponce
	json.Unmarshal(data, &creatResp)
	defer storage.DeletedLocal(creatResp.Metadata.Id)
	err := a.DeleteBin(creatResp.Metadata.Id)
	if err != nil {
		t.Errorf("Ожидалось удачное удаление, получаем ошибку: %v", err)
	}
	defer a.DeleteBin(creatResp.Metadata.Id)
}
