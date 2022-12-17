/*
Задание 2. Анонимные функции
Что нужно сделать
Напишите функцию, которая на вход принимает функцию вида A func (int, int) int, а внутри оборачивает и вызывает её при выходе (через defer).

Вызовите эту функцию с тремя разными анонимными функциями A. Тела функций могут быть любыми, но главное, чтобы все три выполняли разное действие.

Что оценивается
Правильность размеров исходных и конечной матриц.
*/
package main

import "fmt"

const (
	rows       = 3
	RowsToCols = 5
	cols       = 4
)

func wrapper(x, y int, F func(int, int) int) {
	defer fmt.Println(F(x, y))
	fmt.Println("Wrapper START")
}

func main() {
	f_sum := func(x, y int) int { return x + y }
	f_mul := func(x, y int) int { return x * y }
	f_div := func(x, y int) int { return int(x / y) }

	wrapper(1, 7, f_sum)
	wrapper(1, 1, f_mul)
	wrapper(4, 2, f_div)
}
