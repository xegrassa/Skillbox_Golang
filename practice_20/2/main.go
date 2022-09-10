/*
Задание 2. Умножение матриц
Что нужно сделать
Напишите функцию, умножающую две матрицы размерами 3 × 5 и 5 × 4.
*/
package main

import "fmt"

const (
	rows       = 3
	RowsToCols = 5
	cols       = 4
)

func multiplicationM(m1 [rows][RowsToCols]int, m2 [RowsToCols][cols]int) (m3 [rows][cols]int) {
	var row, col int
	for i := 0; i < rows*cols; i++ {
		row = i / cols
		col = i % cols
		for j := 0; j < RowsToCols; j++ {
			m3[row][col] += m1[row][j] * m2[j][col]
		}
	}
	return m3
}

func main() {
	matrix1 := [rows][RowsToCols]int{
		{7, 10, 6, 12, 25},
		{2, 15, 8, 10, 11},
		{3, 15, 10, 45, 1},
	}
	matrix2 := [RowsToCols][cols]int{
		{7, 10, 6, 12},
		{2, 15, 8, 10},
		{3, 15, 10, 45},
		{5, 2, 74, 4},
		{41, 5, 9, 7},
	}
	multiMatrix := multiplicationM(matrix1, matrix2)
	fmt.Println(multiMatrix)
	fmt.Println("===================")

	for i := 0; i < len(multiMatrix); i++ {
		fmt.Println(multiMatrix[i])
	}
}
