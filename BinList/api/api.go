package api

import (
	"BinList/app/bins"
	"BinList/app/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Api struct {
	Bin    *bins.Bin
	keyEnv *config.Config
}
type List struct {
	Name []string `json:"name"`
	Id   []string `json:"ID"`
}

func (api Api) CreatedBin(status string) {
	data, errJs := json.Marshal(api)
	if errJs != nil {
		fmt.Println(errJs)
		return
	}
	reque, errReque := http.NewRequest("PUT", "https://api.jsonbin.io/v3", bytes.NewBuffer(data))
	if errReque != nil {
		fmt.Println(errReque)
		return
	}
	reque.Header.Set("X-Master-Key", api.keyEnv.MasterKey)
	reque.Header.Set("X-Access-Key", api.keyEnv.MasterKey)
	reque.Header.Set("Content-Type", "application/json")
	reque.Header.Set("X-Bin-Private", status)
	_, ok := sendingRequest(reque)
	if ok {
		fmt.Println("Запись прошла успешно")
	}
}
func (api Api) GetBin(id string) {
	reque, errReque := http.NewRequest("GET", "https://api.jsonbin.io/v3", nil)
	if errReque != nil {
		fmt.Println(errReque)
		return
	}
	reque.Header.Set("X-Master-Key", api.keyEnv.MasterKey)
	reque.Header.Set("X-Access-Key", api.keyEnv.AccessKey)
	data, _ := sendingRequest(reque)
	errJs := json.Unmarshal(data, &api.Bin)
	if errJs != nil {
		fmt.Println(errJs)
	}
	fmt.Printf("Name: %s\nID: %s\nPrivate: %t\ntime of creation: %s\n", api.Bin.Name, api.Bin.Id, api.Bin.Private, api.Bin.CreatedAt)
}
func (api Api) ListBin() {
	reque, errReque := http.NewRequest("GET", "https://api.jsonbin.io/v3", nil)
	if errReque != nil {
		fmt.Println(errReque)
		return
	}
	reque.Header.Set("X-Master-Key", api.keyEnv.MasterKey)
	reque.Header.Set("X-Access-Key", api.keyEnv.AccessKey)
	data, _ := sendingRequest(reque)
	var list List
	errJs := json.Unmarshal(data, &list)
	if errJs != nil {
		fmt.Println(errJs)
	}
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
	data, errRead := io.ReadAll(reque.Body)
	if errRead != nil {
		fmt.Println(errRead)
		return nil, false
	}
	return data, true
}
