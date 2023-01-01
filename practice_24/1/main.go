/*
Задание 1. Чётные и нечётные
Что нужно сделать
Напишите функцию, которая принимает массив чисел, а возвращает два массива: один из чётных чисел, второй из нечётных.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func splitArrayByEven(arr [size]int) (even []int, odd []int) {

	for _, v := range arr {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	return
}

func main() {
	mas := [size]int{}

	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		randNum := rand.Intn(10)
		mas[i] = randNum
	}
	fmt.Println("Изначальынй массив: ", mas)

	even, odd := splitArrayByEven(mas)
	fmt.Println("Четный массив: ", even)
	fmt.Println("Нечетный массив: ", odd)

}
