package main

import "fmt"

func main() {

	number1 := 5
	number2 := 10
	text := "Привет, земляне!"

	score := 5
	fmt.Println("Счёт:", score)

	fmt.Println("До вызова функции")
	fmt.Println("")
	hello()
	fmt.Println("После вызова функции")
	fmt.Println("")
	aboba()

	fmt.Println("Вызываю функцию square(5)")
	square(5)
	fmt.Println("Закончил выполнение square(5)")
	fmt.Println("")

	fmt.Println("Вызываю функцию square(10)")
	square(10)
	fmt.Println("Закончил выполнение square(10)")

	fmt.Println("")

	fmt.Println("Вызываю функцию square(3)")
	square(3)
	fmt.Println("Закончил выполнение square(3)")
	fmt.Println("")

	officiantWork("Иван")
	officiantWork("Олег")
	officiantWork("Василий")

	sum(1, 2, "Приветствую вас!")
	sum(10, 33, "Как дела друзья?")
	sum(number1, number2, text)
}

// Допустим мы работаем официантом в ресторане
// и у нас задача:
// для каждого нового гостя нужно:
// 1. накрыть ему стол
// 2. поприветсвовать по имени
// 3. принять заказ
// 4. принести блюдо

func sum(a int, b int, greeting string) {
	s := a + b

	fmt.Println(greeting)

	fmt.Println("Сумма а + б =", s)
	fmt.Println("")
}

func officiantWork(name string) {
	fmt.Println("Накрываю на стол!")
	fmt.Println("Приветствую, господин", name, "!")
	fmt.Println("Я принял заказ!")
	fmt.Println("Я принёс блюдо!")
	fmt.Println("Не хотите ли вы закзаать что-нибудь еще?")
	fmt.Println("Хорошего дня!")
	fmt.Println("")
}

func hello() {

	fmt.Println("Я функция")
	fmt.Println("Меня вызвали!")
	fmt.Println("Я завершаюсь")
	fmt.Println("")
}

func aboba() {

	a := 1
	b := 10
	c := 3

	fmt.Println("Сумма", a+b+c)
	fmt.Println("")
}

func square(x int) {

	fmt.Println("Мы приняли в функции переменную х:", x)
	fmt.Println("x в квадрате:", x*x)

}
