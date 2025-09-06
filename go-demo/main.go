package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTpower = 2
	var userHeight float64
	var userWeight float64
	fmt.Println("___ Калькулятор индекса массы ___")
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userWeight)
	IMT := (userWeight) / math.Pow (userHeight / 100, IMTpower)
	
}
func outputResult(imt float64) {
result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)
}