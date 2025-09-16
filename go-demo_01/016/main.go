package main

import "fmt"

func main() {

	number := 999

	fmt.Println("До", number)

	for i := 1; i <= 10; i += 2 {
		score := 5

		fmt.Println("Итерация №", i)
		fmt.Println("Score", score)
		fmt.Println("")

		score = score + 3
		fmt.Println("Score", score)
		number++
	}

	fmt.Println("После", number)
}
