/*
Задание 2. Анонимные функции
Что нужно сделать
Напишите анонимную функцию, которая на вход получает массив типа integer, сортирует его пузырьком и переворачивает (либо сразу сортирует в обратном порядке, как посчитаете нужным).

Советы и рекомендации
Не забудьте проверить, что вы получили больше чем ноль аргументов.
Подход не важен, можно переписать сортировку пузырьком или отсортировать, а потом перевернуть.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func main() {
	mas := [size]int{}

	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		randNum := rand.Intn(10)
		mas[i] = randNum
	}
	fmt.Println(mas)

	inverseSort := func(mas [size]int) [size]int {
		var isSorted bool

		for i := 0; 0 < size-1-i; i++ {
			isSorted = true
			for j := 0; j < size-1-i; j++ {
				if mas[j] < mas[j+1] {
					mas[j], mas[j+1] = mas[j+1], mas[j]
					isSorted = false
				}
			}

			if isSorted {
				break
			}
		}
		return mas
	}

	fmt.Println(inverseSort(mas))
}
