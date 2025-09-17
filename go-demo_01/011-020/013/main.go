package main

import "fmt"

func main() {

	score := 5

	if score > 6 && score < 15 {

		fmt.Println("Ты попал в яблочко!")

	} else {

		fmt.Println("Ты не попал в яблочко")

	}

	score = 21

	if score < 6 || score > 16 {

		fmt.Println("Ты попал в зону")

	} else {

		fmt.Println("Ты не попал в зону")

	}

}
