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

	/*
	   import (
	   	"fmt"
	   	//"time"
	   	//"math/rand"
	   	)

	   func main() {

	   	for i := 1; i <= 5; i ++ {
	   	fmt.Println("Итерация", i)

	   	if i ==3{
	   		break
	   		}
	   	}

	*/
	/*
	   import ("fmt"
	   	"time"
	   	"math/rand")

	   func main() {


	   //🐤
	   //🟩
	   //🟥

	   	score := 0

	   	fmt.Println("Get ready!")
	   	fmt.Println("Счёт:", score)
	   	fmt.Println("")


	   	for {

	   	fmt.Println("------------------------------")
	   	fmt.Println("")
	   	fmt.Println("Подлетаю к трубе!")
	   		fmt.Println("🐤🟩🟩")
	   	fmt.Println("")

	   	fmt.Println("Пролетаю через трубу!")
	   	fmt.Println("🟩🐤🟩")
	   	fmt.Println("")

	   	if rand.Intn(8) == 1 {
	   		fmt.Println("Я врезался в трубу :(")
	   		fmt.Println("🟩🟥🟩")
	   		fmt.Println("")
	   		break
	   	}

	   	fmt.Println("Пролетел через трубу!")
	   	fmt.Println("🟩 🟩 🐤")
	   	fmt.Println("")

	   	score++
	   	fmt.Println("Счёт:", score)
	   	fmt.Println("------------------------------")
	   	fmt.Println("")
	   	time.Sleep(1000 * time.Millisecond)
	   	}

	   	fmt.Println("Game over!")
	   	fmt.Println("Ваш итоговый счёт", score)
	*/

	/*
	   import (
	   	"fmt"
	   	"time"
	   )

	   func main() {

	   	for i := 1; i <= 5; i++ {
	   		fmt.Println("Новая итерация! i:", i)
	   		time.Sleep(1000 * time.Millisecond)
	   		}
	*/
	/*
	   import (
	   	"fmt"
	   	"math/rand"
	   )

	   	// "math/rand"
	   	// "time"
	   	// 1. Бесконечные циклы
	   	// 2. Ключевое слово break

	   func main() {

	   	fmt.Println("Hello World!")
	   	fmt.Println("Рандомная цифра", rand.Intn(10))
	   	fmt.Println("")

	   	// 0 1
	   	if rand.Intn(4) == 1 {
	   		fmt.Println("Я сгенерировал рандомно число один!")
	   		} else {
	   			fmt.Println("Я не попал в единицу(")
	   			}
	*/

	/*package main

	  import (
	  	"fmt"
	  )


	  func main() {

	  fmt.Println("Начинаю генерацию уровня")
	  fmt.Println("")

	  //🐤
	  //🟩
	  //🟥

	  	for i:=1; i <=5; i++ {

	  	fmt.Println("Труба №", i)
	  	fmt.Println("----")
	  	if i % 2 == 0 {
	  		fmt.Println("🟥🟥")
	  	} else {
	  	fmt.Println("🟩🟩")
	  	}
	  	fmt.Println("----")


	  	fmt.Println("")
	  	}

	  	fmt.Println("Генерация уровня закончена")
	*/
	/*
	   score := 0

	   fmt.Println("Get Ready")
	   fmt.Println("Счёт:", score)
	   fmt.Println("")

	   //🐤
	   //🟩

	   	for i:=1; i <=5; i++ {

	   	fmt.Println("------------------------------")
	   	fmt.Println("Вы подлетаете к трубе!", i)
	   	fmt.Println("🐤 🟩 🟩")
	   	fmt.Println("")

	   	fmt.Println("Вы пролетаете через трубу!", i)
	   	fmt.Println("🟩🐤🟩")
	   	fmt.Println("")

	   	fmt.Println("Вы пролетели через трубу!", i)
	   	fmt.Println("🟩 🟩 🐤")
	   	fmt.Println("")

	   	score++

	   	fmt.Println("Счёт:", score)
	   	fmt.Println("")
	   	}
	*/

	/*
		number:=999

		fmt.Println("До", number)

		for i:=1; i <=10; i+=2 {
			score:=5

			fmt.Println("Итерация №", i)
			fmt.Println("Score", score)
			fmt.Println("")

			score = score + 3

			number++
			}

		fmt.Println("После", number)
	*/

	/*
		fmt.Println("Hello, World!")

		fmt.Println("Hello, World!")

		subscribed := false

		fmt.Println("До условия", subscribed)

		if !subscribed {

			fmt.Println("Я вижу ты не подписан. Подпишись пожалуйста!")

		}

		//	if subscribed {

		//

		//	} else {

		//		fmt.Println("Я вижу ты не подписан. Подпишись пожалуйста!")

		//		}

		fmt.Println("после условия", subscribed)
	*/
	/*	sunny := true


		weekend := false





		if sunny && weekend{


			fmt.Println("Я иду гулять")


			} else {


				fmt.Println("Я сопротивляюсь")


			}





		computerClub := true


		icecream := false





		if computerClub || icecream {


			fmt.Println("Я иду гулять")


			} else {


				fmt.Println("Я сопротивляюсь")


			} */

	/*	score := 5


			if score > 6 && score < 15 {


				fmt.Println("Ты попал в яблочко!")


				} else {


					fmt.Println("Ты не попал в яблочко")


					}




















		//	score := 21


			if score <6 || score > 16{


			fmt.Println("Ты попал в зону")


			} else {


				fmt.Println("Ты не попал в зону")


				}


	*/

	/*	number := 15





		ravno5 := number == 5


		bolshe12 := number > 12





		fmt.Println(ravno5)


		fmt.Println(bolshe12)*/

	/*score := 50





	if score == 12 {


		fmt.Println("Дюжина!")


	} else if score == 21{


		fmt.Println("Очко")


		} else if score == 50 {


			fmt.Println("Полтинник!")


			} else {


			fmt.Println("Ты не попал ни в какую категорию")


	}*/

	/*score := 16


	if score > 15 {


		fmt.Println("Ты мега-красавчик!")


	} else if score > 10 {


		fmt.Println("Ты красавчик!")


	} else {


		fmt.Println("Тебе нужно еще многому научиться")





	}





	fmt.Println("Счёт: ", score)


	fmt.Println("Спасибо за использование нашей игры! Возвращайтесь еще!") */

	/*


		score := 10





		if score > 10 {


				if score > 15 {


					fmt.Println("Ты мега-красавчик!")


				} else {


					fmt.Println("Ты красавчик!")


				}


			} else {


				fmt.Println("Тебе нужно еще многому научиться")


			}





			fmt.Println("Счёт: ", score)


			fmt.Println("Спасибо за использование нашей игры! Возвращайтесь еще!")


	*/

	/*	number := 7





		number ++





		number --





		fmt.Println("number:", number)





		text := "hello"


		text += " world"





		fmt.Println("text: ", text) */

	/*	var number int


		var text string


		var drob float64


		var boolean bool





		fmt.Println("number:", number)


		fmt.Println("text:", text)


		fmt.Println("drob:", drob)


		fmt.Println("boolean:", boolean) */

	/*	var number int


		var text string





		fmt.Println("number", number)


		fmt.Println("text", text)





		number = 15


		text = "Hello world"





		fmt.Println("number", number)


		fmt.Println("text", text) */

	/*	score := 12


		s := "hello"





		var number int = 10


		var text string = "hi" */

	/*	text := "hello"


		text = text + " world"


		fmt.Println(text) */

	/*	score := 11


			celoe := score / 3


			ostatok := score % 3





		//		score = score + 10


		//		score = score + 20


		//		score = score - 15


		//		score = score / 3


		//		score = score * 10





			fmt.Println("Счёт: ", score, celoe, ostatok) */

	/*


		text :="hello"


		score := 0


		drob := 0.5





		fmt.Println(text)


		fmt.Println(score)


		fmt.Println(drob)





		subscribed := false





		fmt.Println("Подписан на Мистера Биста?", subscribed)


	*/

	/*


	   text1 := "Get Ready"


	   text2 := "Game Over"


	   score := 0





	   fmt.Println(text1, score)


	   fmt.Println("Ваш счёт", score)


	   score = score + 1





	   fmt.Println("Вы пролетели через трубу", score)


	   fmt.Println("Ваш счёт", score)


	   score = score + 1


	   fmt.Println("Вы пролетели через трубу", score)





	   fmt.Println("Ваш счёт", score)


	   score = score + 1


	   fmt.Println("Вы пролетели через трубу", score)





	   fmt.Println("Врезались в трубу")


	   fmt.Println(text2)


	*/

}
