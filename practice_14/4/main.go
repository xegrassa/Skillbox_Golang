//Задание 4. Область видимости переменных
//Что нужно сделать
//Напишите программу, в которой будет три функции, попарно использующие разные глобальные переменные. Функции должны
//прибавлять к поданному на вход числу глобальную переменную и возвращать результат. Затем вызовите по очереди три
//функции, передавая результат из одной в другую.

package main

import "fmt"

var (
	global1 = 1
	global2 = 2
	global3 = 3
)

func f1(n int) int {
	return n + global2 + global3
}

func f2(n int) int {
	return n + global1 + global3
}

func f3(n int) int {
	return n + global1 + global2
}

func main() {
	fmt.Println("Введите число:")

	var number int
	_, _ = fmt.Scan(&number)

	number = f3(f2(f1(number)))
	fmt.Printf("Запустив функции по очереди получим: %v\n", number)
}
