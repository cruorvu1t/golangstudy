package main

import "fmt"

func main() {

	text1 := "Get Ready"

	text2 := "Game Over"

	score := 0

	fmt.Println(text1, score)

	fmt.Println("Ваш счёт", score)

	score = score + 1

	fmt.Println("Вы пролетели через трубу", score)

	fmt.Println("Ваш счёт", score)

	score = score + 1

	fmt.Println("Вы пролетели через трубу", score)

	fmt.Println("Ваш счёт", score)

	score = score + 1

	fmt.Println("Вы пролетели через трубу", score)

	fmt.Println("Врезались в трубу")

	fmt.Println(text2)

}
