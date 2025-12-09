package api

import (
	"BinList/app/bins"
	"BinList/app/config"
	"BinList/app/files"
	"BinList/app/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Api struct {
	keyEnv *config.Config
}

func Newapi() *Api {
	return &Api{
		keyEnv: config.NewConfig(),
	}
}
func (api Api) CreatedBin(name, file string) {
	var status string
	fmt.Println("Укажите статус true - публичный, false - приватный")
	fmt.Scan(&status)
	if status != "true" && status != "false" {
		fmt.Println("статус приватности должен быть true или false")
		return
	}
	Bin, err := bins.NewBinCreate(name)
	if err != nil {
		color.Red(err.Error())
		return
	}
	storages, _, erro := storage.NewStorage(files.NewJsonDb(file))
	if erro != nil {
		color.Red(erro.Error())
		return
	} else {
		color.Green("Запись в файл %s прошла успешно\n", file)
	}
	data, erro := storages.AddStorage(*Bin)
	if erro != nil {
		color.Red(erro.Error())
	}
	errJs := json.Unmarshal(data, &storages)
	if errJs != nil {
		fmt.Println(errJs)
		return
	}
	reque, errReque := http.NewRequest("POST", "https://api.jsonbin.io/v3/b/", bytes.NewBuffer(data))
	if errReque != nil {
		fmt.Println(errReque)
		return
	}
	reque.Header.Set("X-Master-Key", api.keyEnv.MasterKey)
	reque.Header.Set("X-Access-Key", api.keyEnv.MasterKey)
	reque.Header.Set("Content-Type", "application/json")
	reque.Header.Set("X-Bin-Private", status)
	resDatas, _ := sendingRequest(reque)
	var b storage.Storage
	errJss := json.Unmarshal(resDatas, &b)
	if errJss != nil {
		fmt.Println(errJss)
		return
	}
	fmt.Println("Отправка бина прошла успешно")
}
func (api Api) UpdateBin(id, file string) {
	Bin, err := bins.NewBinUpdate(id)
	if err != nil {
		color.Red(err.Error())
		return
	}
	storages, _, erro := storage.NewStorage(files.NewJsonDb(file))
	if erro != nil {
		color.Red(erro.Error())
		return
	} else {
		color.Green("Запись в файл %s прошла успешно\n", file)
	}
	data, erro := storages.AddStorage(*Bin)
	if erro != nil {
		color.Red(erro.Error())
	}
	errJs := json.Unmarshal(data, &storages)
	if errJs != nil {
		fmt.Println(errJs)
		return
	}
	reque, errReque := http.NewRequest("PUT", "https://api.jsonbin.io/v3/b/"+id, bytes.NewBuffer(data))
	if errReque != nil {
		fmt.Println(errReque)
		return
	}
	reque.Header.Set("X-Master-Key", api.keyEnv.MasterKey)
	reque.Header.Set("X-Access-Key", api.keyEnv.MasterKey)
	reque.Header.Set("Content-Type", "application/json")
	_, ok := sendingRequest(reque)
	if ok {
		color.Green("Добавление в файл: %s прошла успешно", file)
	}
}
func (api Api) DeleteBin(id string) {
	reque, errReque := http.NewRequest("DELETE", "https://api.jsonbin.io/v3/b/"+id, nil)
	if errReque != nil {
		fmt.Println(errReque)
	}
	reque.Header.Set("X-Master-Key", api.keyEnv.MasterKey)
	reque.Header.Set("X-Access-Key", api.keyEnv.MasterKey)
	_, ok := sendingRequest(reque)
	if ok {
		color.Green("Удление прошло успешно")
	}
}
func (api Api) GetBin(id string) {
	reque, errReque := http.NewRequest("GET", "https://api.jsonbin.io/v3/b/"+id, nil)
	if errReque != nil {
		fmt.Println(errReque)
		return
	}
	reque.Header.Set("X-Master-Key", api.keyEnv.MasterKey)
	reque.Header.Set("X-Access-Key", api.keyEnv.AccessKey)
	data, _ := sendingRequest(reque)
	var stor storage.Storage
	errJs := json.Unmarshal(data, &stor)
	if errJs != nil {
		fmt.Println(errJs)
	}
	fmt.Println(stor)
}

func sendingRequest(reque *http.Request) ([]byte, bool) {
	client := &http.Client{}
	resp, errResp := client.Do(reque)
	if errResp != nil {
		fmt.Println(errResp)
		return nil, false
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println("Error code:", resp.StatusCode)
		return nil, false
	}
	data, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		fmt.Println(errRead)
		return nil, false
	}
	return data, true
}
func ListBin() error {
	allFiles, errDier := os.ReadDir(".")
	if errDier != nil {
		return errDier
	}
	var sliceFile []string
	for _, value := range allFiles {
		if strings.HasSuffix(value.Name(), ".json") {
			sliceFile = append(sliceFile, value.Name())
		}
	}
	var sliceList []storage.Storage
	for _, file := range sliceFile {
		data, errRead := os.ReadFile(file)
		if errRead != nil {
			return errRead
		}
		var listen storage.Storage
		errJs := json.Unmarshal(data, &listen)
		if errJs != nil {
			return errJs
		}
		sliceList = append(sliceList, listen)
	}
	var sliceBin []bins.Bin
	for _, store := range sliceList {
		binss := store.Bins
		sliceBin = append(sliceBin, binss...)
	}
	for _, bin := range sliceBin {
		fmt.Printf("name: %s | ID: %s\n", bin.Name, bin.Id)
	}
	return nil
}
