package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//🐤
	//🟩
	//🟥

	score := 0

	fmt.Println("Get ready!")
	fmt.Println("Счёт:", score)
	fmt.Println("")

	for {

		fmt.Println("------------------------------")
		fmt.Println("")
		fmt.Println("Подлетаю к трубе!")
		fmt.Println("🐤🟩🟩")
		fmt.Println("")

		fmt.Println("Пролетаю через трубу!")
		fmt.Println("🟩🐤🟩")
		fmt.Println("")

		if rand.Intn(8) == 1 {
			fmt.Println("Я врезался в трубу :(")
			fmt.Println("🟩🟥🟩")
			fmt.Println("")
			break
		}

		fmt.Println("Пролетел через трубу!")
		fmt.Println("🟩 🟩 🐤")
		fmt.Println("")

		score++
		fmt.Println("Счёт:", score)
		fmt.Println("------------------------------")
		fmt.Println("")
		time.Sleep(1000 * time.Millisecond)
	}

	fmt.Println("Game over!")
	fmt.Println("Ваш итоговый счёт", score)
}
