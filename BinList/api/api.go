package api

import (
	"BinList/app/config"
	"BinList/app/storage"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

type Api struct {
	keyEnv *config.Config
}
type Meta struct {
	Id string `json:"id"`
}
type CreatedResponce struct {
	Metadata       Meta         `json:"metadata"`
	Recorder storage.Storage `json:"record"`
}

func Newapi() *Api {
	return &Api{
		keyEnv: config.NewConfig(),
	}
}
func (api Api) CreatedBin(name, file string) error {
	var status, password string
	color.Cyan("Укажите ваш пароль. Если пароль не будет указан, он сгенерируется автоматически из 20 символов")
	fmt.Scanln(&password)
	color.Cyan("Укажите статус true - публичный, false - приватный")
	fmt.Scan(&status)
	if status != "true" && status != "false" {
		return errors.New("статус приватности должен быть true или false")
	}
	data, errLoc := storage.CreateLocal(name, "", password, status, file)
	if errLoc != nil {
		return errLoc
	}
	reque, errReque := http.NewRequest("POST", "https://api.jsonbin.io/v3/b/", bytes.NewBuffer(data))
	if errReque != nil {
		return errReque
	}
	k := Newapi()
	reque.Header.Set("X-Master-Key", k.keyEnv.MasterKey)
	reque.Header.Set("X-Access-Key", k.keyEnv.AccessKey)
	reque.Header.Set("Content-Type", "application/json")
	reque.Header.Set("X-Bin-Private", status)
	_, data, errResp := sendingRequest(reque)
	if errResp != nil {
		return errResp
	}
	var creatResp CreatedResponce
	json.Unmarshal(data, &creatResp)
color.Green("Отправка бина прошла успешно, ваш ID: %s", creatResp.Metadata.Id)
	_, errLoacals := storage.CreateLocal(name, creatResp.Metadata.Id, password, status, file)
	if errLoacals != nil {
		return errLoacals
	}
	return nil
}
func (api Api) UpdateBin(id, file string) error {
	var name, status, password string
	color.Cyan("Укажите имя")
	fmt.Scan(&name)
	color.Cyan("Укажите ваш пароль. Если пароль не будет указан, он сгенерируется автоматически из 20 символов")
	fmt.Scanln(&password)
	color.Cyan("Укажите статус true - публичный, false - приватный")
	fmt.Scan(&status)
	if status != "true" && status != "false" {
		return errors.New("статус приватности должен быть true или false")
	}
	data, errLoc := storage.UpdateLocal(name, id, password, status, file)
	if errLoc != nil {
		return errLoc
	}

	reque, errReque := http.NewRequest("PUT", "https://api.jsonbin.io/v3/b/"+id, bytes.NewBuffer(data))
	if errReque != nil {
		return errReque
	}
	k := Newapi()
	reque.Header.Set("X-Master-Key", k.keyEnv.MasterKey)
	reque.Header.Set("X-Access-Key", k.keyEnv.MasterKey)
	reque.Header.Set("Content-Type", "application/json")
	_, _, errResp := sendingRequest(reque)
	if errResp != nil {
		return errResp
	}

	color.Green("Добавление в файл: %s прошла успешно", file)
	return nil
}
func (api Api) DeleteBin(id string) error {
	reque, errReque := http.NewRequest("DELETE", "https://api.jsonbin.io/v3/b/"+id, nil)
	if errReque != nil {
		return errReque
	}
	k := Newapi()
	reque.Header.Set("X-Master-Key", k.keyEnv.MasterKey)
	reque.Header.Set("X-Access-Key", k.keyEnv.MasterKey)
	_, _, errResp := sendingRequest(reque)
	if errResp != nil {
		return errResp
	}
	color.Green("Удаление записи на внешнем сервисе прошло успешно")
	errDel := storage.DeletedLocal(id)
	if errDel != nil {
		return errDel
	}
	return nil
}
func (api Api) GetBin(id string) error {
	reque, errReque := http.NewRequest("GET", "https://api.jsonbin.io/v3/b/"+id, nil)
	if errReque != nil {
		return errReque
	}
	k := Newapi()
	reque.Header.Set("X-Master-Key", k.keyEnv.MasterKey)
	reque.Header.Set("X-Access-Key", k.keyEnv.AccessKey)
	_, data, errResp := sendingRequest(reque)
	if errResp != nil {
		return errResp
	}
	var rec CreatedResponce
	errJs := json.Unmarshal(data, &rec)
	if errJs != nil {
		return errJs
	}
	for _, value := range rec.Recorder.Bins{
		color.Magenta("Name: %s\nID: %s\nPassword: %s\nPrivate: %s\nFileName: %s\nCreatTime: %s\n\n\n", value.Name, value.Id, value.Password, value.Private, value.LocalFile, value.CreatedAt)
	}
	return nil
}

func sendingRequest(reque *http.Request) (*http.Response, []byte, error) {
	client := &http.Client{}
	resp, errResp := client.Do(reque)
	if errResp != nil {

		return nil, nil, errResp
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, nil, errors.New(fmt.Sprintln("Error code:", resp.StatusCode))
	}
	dataBody, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return nil, nil, errRead
	}
	return resp, dataBody, nil
}
