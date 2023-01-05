/*
Задание 1. Конвейер
Цели задания

	Научиться работать с каналами и горутинами.
	Понять, как должно происходить общение между потоками.

Что нужно сделать

	Реализуйте паттерн-конвейер:

Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
Произведение: следующая горутина умножает квадрат числа на 2.
При вводе «стоп» выполнение программы останавливается.

Советы и рекомендации

	Воспользуйтесь небуферизированными каналами и waitgroup.

Что оценивается
Ввод : 3

Квадрат : 9

Произведение : 18
*/
package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var num string

	intChan1 := make(chan int)
	intChan2 := make(chan int)

	var wg sync.WaitGroup

	for {
		fmt.Print("Ввод : ")
		_, _ = fmt.Scan(&num)
		intNumber, err := strconv.Atoi(num)

		if num == "стоп" {
			close(intChan1)
			break
		}
		if err != nil {
			fmt.Println("Было введено не число !")
			continue
		}

		wg.Add(2)
		go square(&wg, intChan1, intChan2)
		go multiple(&wg, intChan2)

		intChan1 <- intNumber

		wg.Wait()
	}
}

func square(wg *sync.WaitGroup, inpChan, outChan chan int) {
	defer wg.Done()

	num := <-inpChan
	fmt.Printf("Квадрат: %v\n", num*num)
	outChan <- num * num
}

func multiple(wg *sync.WaitGroup, inpChan chan int) {
	defer wg.Done()

	num := <-inpChan
	fmt.Printf("Произведение: %v\n", num*2)
}
