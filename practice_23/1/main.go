/*
Задание 1. Сортировка вставками
Что нужно сделать
Напишите функцию, сортирующую массив длины 10 вставками.

Советы и рекомендации
Алгоритм сортировки доступен на «Википедии».
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func sortInsert(mas *[size]int) {
	for i := 1; i <= size-1; i++ {

		for j := 0; j <= i; j++ {
			if mas[i] < mas[j] {
				mas[j], mas[i] = mas[i], mas[j]
			}
		}
	}
}

func main() {
	mas := [size]int{}

	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		randNum := rand.Intn(10)
		mas[i] = randNum
	}
	fmt.Println(mas)

	sortInsert(&mas)
	fmt.Println(mas)
}
