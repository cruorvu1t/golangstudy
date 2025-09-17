package main

import "fmt"

// Стек отложенных вызовов
func main() {
	fmt.Println("Я мэйн и я начался")
	defer func() {
		fmt.Println("Я мэйн и я закончился")
	}()

	database()
}

func database() {
	// создать подключение
	defer func() {
		// закрыть подключение
	}()
	
	boolean := true
	
	if boolean {
		//
		return
	} else {
	
	}
	
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			//
			return
		} else {
			//
			return
		}
			
	// закрыть подключение
}