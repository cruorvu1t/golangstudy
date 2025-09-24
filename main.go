package main

import (
	"fmt"
	"study/greeting"
	"study/user"
)

func main() {
	fmt.Println("Hello, World!")

	greeting.SayHello()

	greeting.SayBad()

	i := greeting.GiveMeInt()
	fmt.Println("i =", i)

	u := user.NewUser("Данил", 50)
	fmt.Println("Name:", u.GetName())
	fmt.Println("Age:", u.GetAge())
	
	u.SetNewName("Алексей")
	u.SetNewAge(25)

	fmt.Println("Name:", u.GetName())
	fmt.Println("Age:", u.GetAge())
}
