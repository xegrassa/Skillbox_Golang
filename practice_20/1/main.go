/*
Задание 1. Подсчёт определителя
Что нужно сделать
Напишите функцию, вычисляющую определитель матрицы размером 3 × 3.
*/
package main

import "fmt"

func calc_determinant(matrix [3][3]int) int {
	var extendedMatrix [3][5]int

	//Строим расширенную матрицу
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			if j >= 3 {
				extendedMatrix[i][j] = matrix[i][j-3]
			} else {
				extendedMatrix[i][j] = matrix[i][j]
			}
		}
	}

	//Складываем главные диагонали
	result := 0
	for i := 0; i < 3; i++ {
		result += extendedMatrix[0][0+i] * extendedMatrix[1][1+i] * extendedMatrix[2][2+i]
	}
	//Вычитаем побочные диагонали
	for i := 0; i < 3; i++ {
		result -= extendedMatrix[0][2+i] * extendedMatrix[1][1+i] * extendedMatrix[2][0+i]
	}

	return result
}

func main() {
	matrix := [3][3]int{
		{1, 20, 1},
		{6, 2, 9},
		{5, 3, 2},
	}

	fmt.Println(calc_determinant(matrix))
}
