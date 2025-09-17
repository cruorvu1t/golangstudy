package main

import "fmt"

func main() {
	number := 15

	var ptr *int = &number

	fmt.Println("number:", number)
	fmt.Println("ptr:", ptr)

	if ptr != nil {
		fmt.Println("Разыменование:", *ptr)
	} else {
		fmt.Println("Получен nil указатель")
	}
}
