package main

import "fmt"

func main() {

	number := 5
	text := "Привет"

	fmt.Println("number", number)
	fmt.Println("text", text)

	foo(number, text)

	fmt.Println("number", number)
	fmt.Println("text", text)

}

func foo(n int, t string) {
	fmt.Println("n", n)
	fmt.Println("t", t)

	n = 100
	t = "Меня поменяли :("
	fmt.Println("n", n)
	fmt.Println("t", t)
}
