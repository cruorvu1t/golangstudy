package main

import (
	"fmt"
	"math/rand"
)

// "math/rand"
// "time"
// 1. Бесконечные циклы
// 2. Ключевое слово break

func main() {

	fmt.Println("Hello World!")
	fmt.Println("Рандомная цифра", rand.Intn(10))
	fmt.Println("")

	// 0 1
	if rand.Intn(4) == 1 {
		fmt.Println("Я сгенерировал рандомно число один!")
	} else {
		fmt.Println("Я не попал в единицу(")
	}
}
