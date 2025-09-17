package main

import "fmt"

func main() {
	name := "Игорь"
	ptr := &name
	fmt.Println("Имя до изменения:", name)
	changeName(ptr)
	fmt.Println("Имя после изменения:", name)

}

func changeName(s *string) {
	*s = "Иван"
}
