package main

import "fmt"

func main() {

	fmt.Println("Hello, World!")

	fmt.Println("Hello, World!")

	subscribed := false

	fmt.Println("До условия", subscribed)

	if !subscribed {

		fmt.Println("Я вижу ты не подписан. Подпишись пожалуйста!")

	}

	/*if subscribed {

	} else {

		fmt.Println("Я вижу ты не подписан. Подпишись пожалуйста!")

	}*/

	fmt.Println("после условия", subscribed)
}
