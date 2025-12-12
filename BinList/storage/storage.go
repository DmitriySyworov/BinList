package storage

import (
	"BinList/app/bins"
	"BinList/app/files"
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Storage struct {
	Bins     []bins.Bin `json:"Bins"`
	UpdateAt time.Time  `json:"updateAt"`
}

func NewStorage(name, id, password, private, file string) (*Storage, error) {
	bin, errBin := bins.NewBin(name, id, password, private, file)
	if errBin != nil {
		return nil, errBin
	}
	var store []bins.Bin
	store = append(store, *bin)
	return &Storage{
		Bins:     store,
		UpdateAt: time.Now(),
	}, nil
}
func CreateLocal(name, id, password, status, file string) ([]byte, error) {
	store, storeError := NewStorage(name, id, password, status, file)
	if storeError != nil {
		return nil, storeError
	}
	data, errJs := json.Marshal(store)
	if errJs != nil {
		return nil, errJs
	}
	errWrite := files.Write(file, data)
	if errWrite != nil {
		return nil, errWrite
	}
	return data, nil
}
func UpdateLocal(name, id, password, status, file string) ([]byte, error) {
	oldData, errRead := files.Read(file)
	if errRead != nil {
		return nil, errRead
	}
	var Store Storage
	json.Unmarshal(oldData, &Store)
	newStor, errStor := NewStorage(name, id, password, status, file)
	if errStor != nil {
		return nil, errStor
	}
	Store.Bins = append(Store.Bins, newStor.Bins...)
	resStruct := &Storage{
		Bins:     Store.Bins,
		UpdateAt: time.Now(),
	}
	data, errJs := json.Marshal(resStruct)
	if errJs != nil {
		return nil, errJs
	}
	errWrite := files.Write(file, data)
	if errWrite != nil {
		return nil, errWrite
	}
	return data, nil
}
func DeletedLocal(id string) error {
	allFiles, errDir := os.ReadDir(".")
	if errDir != nil {
		return errDir
	}
	var sliceStore []Storage
	for _, files := range allFiles {
		if strings.HasSuffix(files.Name(), ".json") {
			data, errRead := os.ReadFile(files.Name())
			if errRead != nil {
				return errRead
			}
			var store Storage
			errJs := json.Unmarshal(data, &store)
			if errJs != nil {
				return errJs
			}
			sliceStore = append(sliceStore, store)
		}
	}
	for _, store := range sliceStore {
		for _, bin := range store.Bins {
			if bin.Id == id {
				os.Remove(bin.LocalFile)
				color.Green("файл %s успешно удален", bin.LocalFile)
				break
			}
		}
	}
	return nil
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
	var sliceList []Storage
	for _, file := range sliceFile {
		data, errRead := os.ReadFile(file)
		if errRead != nil {
			return errRead
		}
		var listen Storage
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
		color.Magenta("name: %s | ID: %s\n", bin.Name, bin.Id)
	}
	return nil
}
