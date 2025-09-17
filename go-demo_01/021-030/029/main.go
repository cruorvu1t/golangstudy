package main

import "fmt"

// Стек отложенных вызовов
func main() {
	fmt.Println("Я мэйн и я начался")
	defer func() {
		fmt.Println("Я мэйн и я закончился")
	}()

	fmt.Println("Hello1")
	fmt.Println("Hello2")
	fmt.Println("Hello3")
	//hello()
	foo()
}

func hello() {
	fmt.Println("Я хэллоу и я начался")
	defer func() {
		fmt.Println("Я хэллоу и я закончился")
	}()
	fmt.Println("Я функция хэллоу 1")
	fmt.Println("Я функция хэллоу 2")
	fmt.Println("Я функция хэллоу 3")
}

func foo() {
	defer func() {
		fmt.Println("Defer1")
	}()
	defer func() {
		fmt.Println("Defer2")
	}()
	defer func() {
		fmt.Println("Defer3")
	}()

	fmt.Println("FOO")
}
