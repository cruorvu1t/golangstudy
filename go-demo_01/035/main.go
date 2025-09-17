package main

import "fmt"

type User struct {
	Name   string
	Rating float64
}

func (u User) Greeting() {
	fmt.Println("Всем привет!")
	fmt.Println("Меня зовут", u.Name)
	fmt.Println("Мой рейтинг", u.Rating)
	fmt.Println("")
}

func (u User) GoodBye() {
	fmt.Println("Всем пока!")
	fmt.Println("Меня звали", u.Name)
	fmt.Println("Мой рейтинг был:", u.Rating)
	fmt.Println("")
}

func (u *User) RatingUP(rating float64) {
	if u.Rating+rating <= 10.0 {
		u.Rating += rating
		fmt.Println("Я добавил рейтинг")
	} else {
		fmt.Println("Я не прошёл валидацию")
	}
}

/*
func RatingUP(u *User, rating float64) {
	if u.Rating + rating <= 10.0 {
		u.Rating += rating
		fmt.Println("Я добавил рейтинг")
	} else {
		fmt.Println("Я не прошёл валидацию")
	}
}
*/

func main() {

	user := User{
		Name:   "Вася",
		Rating: 6.0,
	}

	//	ptr := &user

	fmt.Println("User до:", user)
	fmt.Println("")

	//	Greeting(user)
	//	GoodBye(user)
	//	user.Greeting()
	//	user.GoodBye()
	user.RatingUP(2.0)

	//	RatingUP(ptr, 2.0)

	fmt.Println("User после:", user)

	//	fmt.Println("User:", user)
	//	fmt.Println("Имя до:", user.Name)
	//	user.Name = "Игорь"
	//	fmt.Println("Имя после:", user.Name)

}
