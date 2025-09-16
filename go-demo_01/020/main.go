package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println("Новая итерация! i:", i)
		time.Sleep(1000 * time.Millisecond)
	}
}
