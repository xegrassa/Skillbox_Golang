/*
Задание 1. Слияние отсортированных массивов
Что нужно сделать
Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
*/
package main

import "fmt"

func main() {
	const (
		maxSizeArr1 = 4
		maxSizeArr2 = 5
		maxSizeArr3 = 9
	)
	arr1 := [maxSizeArr1]int{1, 3, 5, 7}
	arr2 := [maxSizeArr2]int{0, 2, 4, 6, 8}
	var arr3 [maxSizeArr3]int

	var arr1_index, arr2_index int
	for i, _ := range arr3 {

		isArr1IndexEnd := arr1_index == maxSizeArr1
		if isArr1IndexEnd {
			arr3[i] = arr2[arr2_index]
			arr2_index++
			continue
		}

		isArr2IndexEnd := arr2_index == maxSizeArr2
		if isArr2IndexEnd {
			arr3[i] = arr2[arr1_index]
			arr1_index++
			continue
		}

		if arr1[arr1_index] <= arr2[arr2_index] {
			arr3[i] = arr1[arr1_index]
			arr1_index++
		} else {
			arr3[i] = arr2[arr2_index]
			arr2_index++
		}
	}

	fmt.Println(arr3)
}
