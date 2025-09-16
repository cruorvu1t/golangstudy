package main

import "fmt"

func main() {
	sunny := true

	weekend := false

	if sunny && weekend {

		fmt.Println("Я иду гулять")

	} else {

		fmt.Println("Я сопротивляюсь")

	}

	computerClub := true

	icecream := false

	if computerClub || icecream {

		fmt.Println("Я иду гулять")

	} else {

		fmt.Println("Я сопротивляюсь")

	}
}
