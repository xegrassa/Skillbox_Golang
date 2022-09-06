package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxValuePoint = 20

func genPoint() (int, int) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(maxValuePoint)
	y := rand.Intn(maxValuePoint)

	return x, y
}

func changePoint(x, y int) (int, int) {
	x = 2*x + 10
	y = -3*y - 5
	return x, y
}

func main() {
	fmt.Println("Функция, возвращающая несколько значений.")
	var x, y int

	for i := 1; i <= 3; i++ {
		x, y = genPoint()
		fmt.Printf("%v Точка было: x=%v y=%v\n", i, x, y)
		x, y = changePoint(x, y)
		fmt.Printf("%v Точка стало: x=%v y=%v\n", i, x, y)

	}
}

//Напишите программу, которая с помощью функции генерирует три случайные точки в двумерном пространстве
//(две координаты), а затем с помощью другой функции преобразует эти координаты по формулам:
//x = 2 × x + 10, y = −3 × y − 5.
