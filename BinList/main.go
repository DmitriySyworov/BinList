package main

import (
	"BinList/app/bins"
	"BinList/app/files"
	"BinList/app/storage"
	"fmt"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(){
	"1": createdBin,
	"2": readerJson,
	"3": findBins,
}

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		printErr("Не удалось получить переменные окружения. Дальнейшее выполнение программы невозможно!")
	}
	color.Yellow("__Бинлист_Менеджер__")
	for {
		var choice string
		color.Blue(`Укажите вариант выбора:
1. Добавить бин в уже существующий файл/создать новый файл с бинами
2. Прочитать файл JSON
3. Найти определенный бин
4. Выход`)
		fmt.Scan(&choice)
		funcMenu := menu[choice]
		if funcMenu == nil {
			printErr("Конец программы")
			return
		}
		funcMenu()
	}
}
func createdBin() {
	var name, id, private, nameFiles string
	color.Cyan("Укажите имя")
	fmt.Scan(&name)
	color.Cyan("Укажите ID")
	fmt.Scan(&id)
	color.Cyan("Укажите статус приватности: true - публичный, false - приватный")
	fmt.Scan(&private)
	Bin, err := bins.NewBin(name, id, private)
	if err != nil {
		printErr(err)
		return
	}
	color.Cyan("Укажите название файла (Обязательно в формате JSON)")
	fmt.Scan(&nameFiles)
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
func findBins() {
	var searchStr, searchChoice, nameFiles string
	color.Cyan("Укажите название файла json, где находится бин, который вы ищете")
	fmt.Scan(&nameFiles)
	color.Cyan("Укажите каким образом вы хотите найти Бин по name или ID?")
	fmt.Scan(&searchChoice)
	storages, _, err := storage.NewStorage(files.NewJsonDb(nameFiles))
	if err != nil {
		printErr(err)
		return
	}
	switch searchChoice {
	case "name":
		color.Cyan("Укажите точное имя бина")
		fmt.Scan(&searchStr)
		foundBin, err := storages.FindBin(searchStr, func(bin *bins.Bin, str string) bool { return bin.Name == str })
		SaveFind(foundBin, err)
	case "ID":
		color.Cyan("Укажите точное ID бина")
		fmt.Scan(&searchStr)
		foundBin, err := storages.FindBin(searchStr, func(bin *bins.Bin, str string) bool { return bin.Id == str })
		SaveFind(foundBin, err)
	default:
		printErr("Такого варианта выбора не существует")
		return
	}
}
func SaveFind(findBin []bins.Bin, err error) {
	if err != nil {
		printErr(err)
	}
	for _, b := range findBin {
		color.Magenta("Name: %s\nID: %s\nPrivate: %t\ntime of creation: %s\n", b.Name, b.Id, b.Private, b.CreatedAt)
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
