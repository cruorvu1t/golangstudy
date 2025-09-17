package main

import "fmt"

func main() {
	s := str()

	fmt.Println(s)
}

func str() string {
	subscribed := false

	if subscribed {
		return "Ты подписан на канал!"
	} else {
		return "Ты НЕ подписан на канал!"
	}
	// return "Это моя строка!"
}
