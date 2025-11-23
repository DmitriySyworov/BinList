package main

import "fmt"

var menu = map[string]func(){
	"1": conversionEURonUSD,
	"2": conversionEURonRUB,
	"3": conversionUSDonEUR,
	"4": conversionUSDonRUB,
	"5": conversionRUBonEUR,
	"6": conversionRUBonUSD,
}

const EU = 1.17
const ER = 97.35
const UE = 0.85
const UR = 82.9
const RE = 0.01
const RU = 0.012

func main() {
	fmt.Println("__Конвертор валют__")
	var choiceOperation string
	for {
		fmt.Println(`Выберите операцию:
1. Конвертировать EUR в USD
2. Конвертировать EUR в RUB
3. Конвертировать USD в EUR
4. Конвертировать USD в RUB
5. Конвертировать RUB в EUR
6. Конвертировать RUB в USD
7. Выход`)
		fmt.Scan(&choiceOperation)
		funcMenu := menu[choiceOperation]
		if funcMenu == nil {
			fmt.Println("Выход")
			return
		}
		funcMenu()
	}
}
func conversionEURonUSD() {
	var quantity float64
	fmt.Println("Укажите количество евро для конвертации их в доллары")
	fmt.Scan(&quantity)
	fmt.Printf("%.2f EUR = %.2f USD\n", quantity, EU*quantity)
}
func conversionEURonRUB() {
	var quantity float64
	fmt.Println("Укажите количество евро для конвертации их в рубли")
	fmt.Scan(&quantity)
	fmt.Printf("%.2f EUR = %.2f RUB\n", quantity, ER*quantity)
}
func conversionUSDonEUR() {
	var quantity float64
	fmt.Println("Укажите количество долларов для конвертации их в евро")
	fmt.Scan(&quantity)
	fmt.Printf("%.2f USD = %.2f EUR\n", quantity, UE*quantity)
}
func conversionUSDonRUB() {
	var quantity float64
	fmt.Println("Укажите количество долларов для конвертации их в рубли")
	fmt.Scan(&quantity)
	fmt.Printf("%.2f USD = %.2f RUB\n", quantity, UR*quantity)
}
func conversionRUBonEUR() {
	var quantity float64
	fmt.Println("Укажите количество рублей для конвертации их в евро")
	fmt.Scan(&quantity)
	fmt.Printf("%.2f RUB = %.2f EUR\n", quantity, RE*quantity)
}
func conversionRUBonUSD() {
	var quantity float64
	fmt.Println("Укажите количество рублей для конвертации их в доллары")
	fmt.Scan(&quantity)
	fmt.Printf("%.2f RUB = %.2f USD\n", quantity, RU*quantity)
}
