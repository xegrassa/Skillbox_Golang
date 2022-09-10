/*
Задание 2. Сортировка пузырьком
Что нужно сделать
Отсортируйте массив длиной шесть пузырьком.
*/
package main

import "fmt"

func main() {
	const arrSize = 6
	var arr [arrSize]int

	fmt.Println("Введите элементы массива")
	for i := 0; i < arrSize; i++ {
		fmt.Printf("Введите %v элемент массива:", i+1)
		_, _ = fmt.Scan(&arr[i])
	}
	fmt.Println("Введеный массив")
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println("Отсортированный массив")
	fmt.Println(arr)
}
