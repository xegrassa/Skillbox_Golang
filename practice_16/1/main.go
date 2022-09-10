/*
Что нужно сделать
В одном из модулей мы учились выводить шахматное поле в консоль, используя циклы.
Произведите рефакторинг одного из вариантов решения этого задания.
Выведите в функцию процесс вывода шахматного поля в консоль.
Постарайтесь подобрать говорящие имена для функции и переменных вместо a, b, k и j.
*/
package main

import "fmt"

func main() {
	fmt.Println("Введите ширину и высоту поля")

	var width, height int
	_, _ = fmt.Scan(&width, &height)

	for rowIndex := 0; rowIndex < height; rowIndex++ {
		for colIndex := 0; colIndex < width; colIndex++ {
			printChessCell(rowIndex, colIndex)
		}
		fmt.Println()
	}
	fmt.Println("end")
}

func printChessCell(firstIndex int, secondIndex int) {
	if (firstIndex+secondIndex)%2 == 0 {
		fmt.Print("*")
	} else {
		fmt.Print(" ")
	}
}
