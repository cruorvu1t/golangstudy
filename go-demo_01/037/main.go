package main

import "fmt"

type User struct {
	// Имя пользователя
	// Правило: имя должно быть не пустое
	Name string // ""

	// Возраст пользователя
	// Правило: возраст должен быть больше нуля и меньше 150
	Age int // 0

	// Номер телефона
	// Правило: не должен быть пустым
	PhoneNumber string // ""

	// Закрыт ли профиль
	IsClose bool // false

	// Рейтинг пользователя
	// Правило: должен быть больше или равен нолу и меньше или равен десяти
	Rating float64 // 0.0 -- 10.0
}

// Конструктор NewUser - создание нового пользователя с валидацией
func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) User {
	fmt.Println("Валидирую имя")
	if name == "" {
		fmt.Println("Имя не прошло валидацию!")
		return User{}
	}

	fmt.Println("Валидирую возраст")
	if age <= 0 || age >= 150 {
		fmt.Println("Возраст не прошел валидацию!")
		return User{}
	}

	fmt.Println("Валидирую номер телефона")
	if phoneNumber == "" {
		fmt.Println("Номер телефона не прошел валидацию!")
		return User{}
	}

	fmt.Println("Валидирую рейтинг")
	if rating < 0.0 || rating > 10.0 {
		fmt.Println("Рейтинг не прошел валидацию!")
		return User{}
	}

	return User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		IsClose:     isClose,
		Rating:      rating,
	}
}

func (u *User) ChangeName(newName string) {
	if newName != "" {
		u.Name = newName
	}
}

func (u *User) ChangeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}

func (u *User) ChangePhoneNumber(newPhoneNumber string) {
	if newPhoneNumber != "" {
		u.PhoneNumber = newPhoneNumber
	}
}

func (u *User) CloseAccount() {
	u.IsClose = true
}

func (u *User) OpenAccount() {
	u.IsClose = false
}

func (u *User) RatingUP(rating float64) {
	if u.Rating+rating <= 10.0 {
		u.Rating += rating
	}
}

func (u *User) RatingDOWN(rating float64) {
	if u.Rating-rating >= 0.0 {
		u.Rating -= rating
	}
}

func main() {

	user := NewUser(
		"Данил",
		50,
		"+8 888 888 88 88",
		false,
		4.5,
	)

	fmt.Println("User:", user)

	fmt.Println("Рейтинг до:", user.Rating)

	user.RatingUP(2.4)

	fmt.Println("Рейтинг после:", user.Rating)
}
