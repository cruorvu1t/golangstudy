package main

import "fmt"

func main() {
	fmt.Println("До вызова функции")
	fmt.Println("")

	text1 := "Hello "
	text2 := "World"
	fmt.Println("После вызова функции")
	fmt.Println("")

	text3 := sumString(text1, text2)

	fmt.Println("text3:", text3)
}

func sumString(s1 string, s2 string) string {
	s := s1 + s2

	return s
}
