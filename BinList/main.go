package main

import (
	"BinList/app/bins"
	"BinList/app/storage"
	"fmt"
)

func main() {
	getMenu()
}
func getMenu() {
	fmt.Println("__Бинлист_Менеджер__")
exit:
	for {
		var choice int
		fmt.Println(`Укажите вариант выбора:
1. Добавить бин в уже существующий файл/создать новый файл с бинами
2. Прочитать файл JSON
3. Выход`)
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var name, id, private string
			fmt.Println("Укажите имя")
			fmt.Scan(&name)
			fmt.Println("Укажите ID")
			fmt.Scan(&id)
			fmt.Println("Укажите статус приватности: true - публичный, false - приватный")
			fmt.Scan(&private)
			Bin, err := bins.NewBin(name, id, private)
			if err != nil {
				fmt.Println(err)
				break exit
			}
			vault, nameFile, erro := storage.NewStorage()
			if erro != nil {
				fmt.Println(erro)
				break exit
			}
			vault.AddStorage(*Bin, nameFile)
		case 2:
			bin, err3 := storage.NewReadFile()
			if err3 != nil {
				fmt.Println("Такого файла не существует, либо же вы неправильно указали формат")
				break exit
			}
			for _, value := range bin.Bins {
				value.OutputBins()
			}
		case 3:
			fmt.Println("Конец программы")
			break exit
		default:
			fmt.Println("Такого варианта выбора не существует, пожалуйста повторите ввод")
			continue

		}
	}
}
