package main

import (
	"BinList/app/api"
	"BinList/app/bins"
	"BinList/app/files"
	"BinList/app/storage"
	"flag"
	"fmt"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		printErr("Не удалось получить переменные окружения. Дальнейшее выполнение программы невозможно!")
	}
	create := flag.Bool("create", false, "create Bin")
	update := flag.Bool("update", false, "update Bin")
	delete := flag.Bool("delete", false, "delete Bin")
	get := flag.Bool("get", false, "get Bin")
	list := flag.Bool("list", false, "list Bin")
	file := flag.String("file", "", "file name")
	name := flag.String("name", "", "name")
	id := flag.String("id", "", "id")
	flag.Parse()
	switch {
	case *create:
		createdBin(*name, *id, *file)
	case *update:
	case *delete:
	case *get:
		getBins(*id)
	case *list:
	}
}
func createdBin(name, id, nameFiles string) {
	var private string
	resPrivate := false
	color.Cyan("Укажите статус приватности: true - публичный, false - приватный")
	fmt.Scan(&private)
	if private != "true" && private != "false" {
		fmt.Println("Статус приватности должен быть  true  или  false")
		return
	} else if private == "true" {
		resPrivate = true
	}
	var ap api.Api
	ap.CreatedBin(private)
	Bin, err := bins.NewBin(name, id, resPrivate)
	if err != nil {
		printErr(err)
		return
	}
	storages, names, erro := storage.NewStorage(files.NewJsonDb(nameFiles))
	if erro != nil {
		printErr(erro)
		return
	} else {
		color.Green("Запись в файл %s прошла успешно\n", nameFiles)
	}
	erro = storages.AddStorage(*Bin, names)
	if erro != nil {
		printErr(erro)
	}
}
func readerJson() {
	var nameFiles string
	color.Cyan("Укажите название файла (Обязательно в формате JSON)")
	fmt.Scan(&nameFiles)
	bin, err3 := storage.NewReadFile(files.NewJsonDb(nameFiles))
	if err3 != nil {
		printErr("Такого файла не существует, либо же вы неправильно указали формат")
		return
	}
	for _, value := range bin.Bins {
		value.OutputBins()
	}
}
func getBins(id string) {
	var storages storage.StorageWithBd
	foundBin, err := storages.FindBin(id)
	for _, b := range foundBin {
		color.Magenta("Name: %s\nID: %s\nPrivate: %t\ntime of creation: %s\n", b.Name, b.Id, b.Private, b.CreatedAt)
	}
	if err != nil {
		return
	}
}

func printErr(err any) {
	switch v := err.(type) {
	case string:
		color.Red(v)
	case error:
		color.Red(v.Error())
	case int:
		color.Red("Код ошибки:", v)
	default:
		color.Red("Неизвестная ошибка")
	}
}
