package greeting

import "fmt"

func SayHello() {
	fmt.Println("Hello, golang packages!")

	j := GiveMeInt()
	fmt.Println("from sayhello:", j)
}
