package main

import (
	"fmt"
	"math/rand"
	"time"
)

func isEven(a int) bool {
	return a%2 == 0
}

func main() {
	fmt.Println("Функция, возвращающая результат.")

	rand.Seed(time.Now().UnixNano())
	z := rand.Int()

	fmt.Println("Число:", z)
	fmt.Println(isEven(z))

}
