package main

import (
	"BinList/app/bins"
	"BinList/app/files"
	"BinList/app/storage"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	getMenu()
}
func getMenu() {
	fmt.Println("__Бинлист_Менеджер__")
exit:
	for {
		choice := promtData([]string{`Укажите вариант выбора:
1. Добавить бин в уже существующий файл/создать новый файл с бинами
2. Прочитать файл JSON
3. Выход`})
		fmt.Scan(&choice)
		switch choice {
		case "1":
			var name, id, private string
			promtData([]string{"Укажите имя"})
			fmt.Scan(&name)
			promtData([]string{"Укажите ID"})
			fmt.Scan(&id)
			promtData([]string{"Укажите статус приватности: true - публичный, false - приватный"})
			fmt.Scan(&private)
			Bin, err := bins.NewBin(name, id, private)
			if err != nil {
				fmt.Println(err)
				break exit
			}
			var nameFiles string
			promtData([]string{"Укажите название файла, из которого выхотите прочитать бины. (Обязательно в формате JSON)"})
			fmt.Scan(&nameFiles)
			vault, names, erro := storage.NewStorage(files.NewJsonDb(nameFiles))
			if erro != nil {
				fmt.Println(erro)
				break exit
			}
			vault.AddStorage(*Bin, names)
		case "2":
			var nameFiles string
			promtData([]string{"Укажите название файла, из которого выхотите прочитать бины. (Обязательно в формате JSON)"})
			fmt.Scan(&nameFiles)
			bin, err3 := storage.NewReadFile(files.NewJsonDb(nameFiles))
			if err3 != nil {
				fmt.Println("Такого файла не существует, либо же вы неправильно указали формат")
				break exit
			}
			for _, value := range bin.Bins {
				value.OutputBins()
			}
		case "3":
			fmt.Println("Конец программы")
			break exit
		default:
			fmt.Println("Такого варианта выбора не существует, пожалуйста повторите ввод")
			continue

		}
	}
}
func promtData[T any](promt []T) string {
	var usChoice string
	for _, value := range promt {
		color.Cyan("%v", value)
	}
	return usChoice
}
