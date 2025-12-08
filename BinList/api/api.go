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
func (api Api) CreatedBin(name, file string) {
	var status string
	fmt.Println("Укажите статус true - публичный, false - приватный")
	fmt.Scan(&status)
	if status != "true" && status != "false" {
		fmt.Println("статус приватности должен быть true или false")
		return
	}
	bin, errbin := bins.NewBinCreate(name)
	if errbin != nil {
		fmt.Println(errbin)
		return
	}
	_, _, erro := storage.NewStorage(files.NewJsonDb(file))
	if erro != nil {
		color.Red(erro.Error())
		return
	} else {
		color.Green("Запись в файл %s прошла успешно\n", file)
	}
	data, errJs := json.Marshal(bin)
	if errJs != nil {
		fmt.Println(errJs)
		return
	}
	reque, errReque := http.NewRequest("POST", "https://api.jsonbin.io/v3", bytes.NewBuffer(data))
	if errReque != nil {
		fmt.Println(errReque)
		return
	}
	reque.Header.Set("X-Master-Key", api.keyEnv.MasterKey)
	reque.Header.Set("X-Access-Key", api.keyEnv.MasterKey)
	reque.Header.Set("Content-Type", "application/json")
	reque.Header.Set("X-Bin-Private", status)
	resDatas, _:= sendingRequest(reque)
	var b bins.Bin
	errJss := json.Unmarshal(resDatas, &b)
		if errJss != nil{
			fmt.Println(errJss)
			return
		}
		fmt.Println("Отправка бина прошла успешно, ваш id:", b.Id)
}
func (api Api) UpdateBin(id, file string) {
	bin, errbin := bins.NewBinUpdate(id)
	if errbin != nil {
		fmt.Println(errbin)
		return
	}
storages, _, erro := storage.NewStorage(files.NewJsonDb(file))
	if erro != nil {
		color.Red(erro.Error())
		return
	} 
	 erro = storages.AddStorage(*bin)
  if erro != nil {
    color.Red(erro.Error())
  }
  data, errJs := json.Marshal(bin)
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
	var bin bins.Bin
	errJs := json.Unmarshal(data, &bin)
	if errJs != nil {
		fmt.Println(errJs)
	}
	fmt.Printf("Name: %s\nID: %s\ntime of creation: %s\n", bin.Name, bin.Id, bin.CreatedAt)
}

func sendingRequest(reque *http.Request) ([]byte, bool){
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
	data, errRead := io.ReadAll(reque.Body)
	if errRead != nil {
		fmt.Println(errRead)
		return nil, false
	}
	return data, true
}
func ListBin()error{
	allFiles, errDier := os.ReadDir(".")
	if errDier != nil {
		return errDier
	}
	var sliceFile []string
	for _, value := range allFiles{
		if strings.HasSuffix(value.Name(), ".json"){
			sliceFile = append(sliceFile, value.Name())
		}
	}
	var sliceList []storage.Storage
	for _, file := range sliceFile{
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
	for _, store  := range sliceList{
		binss :=  store.Bins
		sliceBin = append(sliceBin, binss...)
	}
	for _, bin := range sliceBin{
		fmt.Printf("name: %s | ID: %s\n",bin.Name, bin.Id)
	}
	return nil
}	

