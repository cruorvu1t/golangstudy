package main

import "fmt"

func main() {
	score := 11

	celoe := score / 3

	ostatok := score % 3

	score = score + 10

	score = score + 20

	score = score - 15

	score = score / 3

	score = score * 10

	fmt.Println("Счёт: ", score, celoe, ostatok)
}
