package main

import "fmt"

// 1. return в функции, которая ничего не возвращает
// 2. глобальные переменные

var number int = 5

// 5
func main() {
	number *= 4
	// 20
	greeting("")

	fmt.Println("number", number)

	kvadrat := square(5)

	fmt.Println("Квадрат:", kvadrat)
}

func greeting(name string) {
	number /= 2
	// 10
	if name == "" {
		fmt.Println("Вы передали пустое имя :(")
		return
	}

	fmt.Println("Привет, уважаемый", name)
}

func square(n int) int {
	s := n * n

	return s
}
