/*
Задание 1. Расчёт по формуле
Что нужно сделать
Напишите функцию, производящую следующие вычисления.

S = 2 × x + y ^ 2 − 3/z, где x — int16, y — uint8, a z — float32.

Тип S должен быть во float32.

Советы и рекомендации
Обратите внимание, к какому типу надо привести конечный результат.
*/
package main

import "fmt"

const (
	rows       = 3
	RowsToCols = 5
	cols       = 4
)

func evaluate(x int16, y uint8, z float32) float32 {
	var s = 2*float32(x) + float32(y)*float32(y) - 3/z
	return s
}

func main() {
	var x int16 = -10
	var y uint8 = 10
	var z float32 = 0.5

	var result float32 = evaluate(x, y, z)
	fmt.Printf("%T - %v", result, result)
}
