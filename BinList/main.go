package main

import (
	"BinList/app/api"
	"BinList/app/storage"
	"flag"
	"fmt"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(color.RedString("Не удалось получить переменные окружения. Дальнейшее выполнение программы невозможно!"))
	}
	menu()
}
func menu() {
	create := flag.Bool("create", false, "create Bin")
	update := flag.Bool("update", false, "update Bin")
	delete := flag.Bool("delete", false, "delete Bin")
	get := flag.Bool("get", false, "get Bin")
	list := flag.Bool("list", false, "list Bin")
	help := flag.Bool("help", false, "helps")
	file := flag.String("file", "", "file name")
	name := flag.String("name", "", "name")
	id := flag.String("id", "", "id")
	flag.Parse()
	var ap api.Api
	switch {
	case *create:
		var status string
		color.Cyan("Укажите статус true - публичный, false - приватный")
		fmt.Scan(&status)
		_, errCreate := ap.CreatedBin(*name, *file, status)
		if errCreate != nil {
			color.Red(errCreate.Error())
			return
		}
	case *update:
		var status string
		color.Cyan("Укажите статус true - публичный, false - приватный")
		fmt.Scan(&status)
		_, errUpdate := ap.UpdateBin(*id, *file, status)
		if errUpdate != nil {
			color.Red(errUpdate.Error())
			return
		}
	case *delete:
		errDelete := ap.DeleteBin(*id)
		if errDelete != nil {
			color.Red(errDelete.Error())
			return
		}
	case *get:
		errGet := ap.GetBin(*id)
		if errGet != nil {
			color.Red(errGet.Error())
			return
		}
	case *list:
		errList := storage.ListBin()
		if errList != nil {
			color.Red(errList.Error())
			return
		}
	case *help:
		color.Yellow(`--create --file=< > --name=< > Чтобы создать новый файл
--update --file=< > --id=< > Чтобы обновить уже существующий файл
--delete --id=< > Чтобы удалить файл 
--get --id=< > Чтобы найти файл
--list Получить список имен и  id бинов`)
	default:
		color.Red("Ошибка: введите команду --help для помощи")
	}
}
