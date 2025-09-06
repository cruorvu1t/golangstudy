package main

import (
	"fmt"
	"math"
)

func main() {
	var userWeight float64
	var userHeight float64
	fmt.Println("___ Калькулятор индекса массы ___")
	userWeight, userHeight = getUserInput(userWeight, userHeight)
	IMT := calculateIMT(userWeight, userHeight)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)
}

func calculateIMT(userWeight float64, userHeight float64) float64{
	const IMTpower = 2
	IMT := userWeight / math.Pow (userHeight / 100, IMTpower)
	return IMT
}

func getUserInput(userWeight float64, userHeight float64) (float64, float64) {
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userWeight)
	return userWeight, userHeight
}
