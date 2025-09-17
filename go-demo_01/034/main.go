package main

import "fmt"

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClose     bool
	Rating      float64
}

func main() {

	user := User{
		Name:        "Сергей",
		Age:         64,
		PhoneNumber: "+7 777 777 77 77",
		IsClose:     true,
		Rating:      1.5,
	}

	fmt.Println("User:", user)
	fmt.Println("Имя до:", user.Name)
	user.Name = "Игорь"
	fmt.Println("Имя после:", user.Name)

}
