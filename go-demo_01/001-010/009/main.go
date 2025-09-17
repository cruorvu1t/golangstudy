package main

import "fmt"

func main() {

	score := 10

	if score > 10 {

		if score > 15 {

			fmt.Println("Ты мега-красавчик!")

		} else {

			fmt.Println("Ты красавчик!")

		}

	} else {

		fmt.Println("Тебе нужно еще многому научиться")

	}

	fmt.Println("Счёт: ", score)

	fmt.Println("Спасибо за использование нашей игры! Возвращайтесь еще!")

}
