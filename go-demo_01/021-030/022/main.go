package main

import (
	"fmt"
	//"time"
	//"math/rand"
)

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println("Итерация", i)

		if i == 3 {
			break
		}
	}

}
