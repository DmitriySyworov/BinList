package main

import (
	"BinList/app/api"
	"flag"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Не удалось получить переменные окружения. Дальнейшее выполнение программы невозможно!")
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
		ap.CreatedBin(*name, *file)
	case *update:
		ap.UpdateBin(*id, *file)
	case *delete:
		ap.DeleteBin(*id)
	case *get:
		ap.GetBin(*id)
	case *list:
		api.ListBin()
	case *help:
		color.Cyan(`--create --file=< > --name=< > Чтобы создать новый файл
--update --file=< > --id=< > Чтобы обновить уже существующий файл
--delete --id=< > Чтобы удалить файл 
--get --id=< > Чтобы найти файл
--list Получить список имен и  id бинов`)
	default:
		color.Red("Ошибка: введите команду --help для помощи")
	}
}
