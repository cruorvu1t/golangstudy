package main

import "fmt"

func main() {
	fmt.Println("До вызова функции")
	fmt.Println("")

	number := sum(1, 2)
	fmt.Println("После вызова функции")
	fmt.Println("")

	fmt.Println("number:", number)
}

func sum(a int, b int) int {
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	s := a + b

	return s
}
