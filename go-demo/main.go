package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTpower = 2
	userHeight := 1.8
	userWeigt := 100.0
	fmt.Println("___ Калькулятор индекса массы ___")
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userWeigt)
	IMT := (userWeigt) / math.Pow (userHeight / 100, IMTpower)
	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
}
