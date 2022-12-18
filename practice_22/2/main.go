/*
Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)
Что нужно сделать
Заполните упорядоченный массив из 12 элементов и введите число.
Необходимо реализовать поиск первого вхождения заданного числа в массив.
Сложность алгоритма должна быть минимальная.

Что оценивается
Верность индекса.

При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе числа 2 программа должна вывести индекс 1.
*/
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const maxLen = 12

func main() {
	mySlice := []int{}

	rand.Seed(time.Now().Unix())
	for i := 0; i < maxLen; i++ {
		randNum := rand.Intn(10)
		mySlice = append(mySlice, randNum)
	}

	sort.Ints(mySlice)
	fmt.Println(mySlice)

	fmt.Print("Введите целое число: ")
	var scannedNum int
	_, _ = fmt.Scan(&scannedNum)

	minIndex := 0
	maxIndex := maxLen - 1

	for minIndex <= maxIndex {
		index := (minIndex + maxIndex) / 2
		foundNum := mySlice[index]

		fmt.Printf("minIndex=%v - maxIndex=%v - index=%v - foundNum=%v\n", minIndex, maxIndex, index, foundNum)

		if foundNum > scannedNum {
			maxIndex = index - 1
			continue
		}

		if foundNum < scannedNum {
			minIndex = index + 1
			continue
		}

		if foundNum == scannedNum {
			for {
				// Чтобы убрать Панику в условии ниже из-за возможности выхода индекса за границу
				if index == 0 {
					fmt.Println(index)
					break
				}

				if mySlice[index-1] != mySlice[index] {
					fmt.Println(index)
					break
				}
				index -= 1
			}
			break
		}
	}
	fmt.Println(-1)
}
