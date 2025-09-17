package main

import "fmt"

func main() {
	fmt.Println("Я мэйн и я начался")
	number := 10

	pointer := &number

	fmt.Println("Это просто адрес в памяти:", pointer)

	foo(pointer)

	//boo(number)

	//fmt.Println(number)
	fmt.Println("Разыменовывание pointer:", *pointer)

}

func foo(n *int) {
	fmt.Println(n)
	*n = 10
	fmt.Println(*n)
}

func boo(n int) {
	n = 10
}
