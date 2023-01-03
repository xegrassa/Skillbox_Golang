/*
Цель задания
Написать программу аналог cat.

Что нужно сделать

	Программа должна получать на вход имена двух файлов, необходимо  конкатенировать их содержимое, используя strings.Join.

Что оценивается

	При получении одного файла на входе программа должна печатать его содержимое на экран.
	При получении двух файлов на входе программа соединяет их и печатает содержимое обоих файлов на экран.
	Если программа запущена командой go run firstFile.txt secondFile.txt resultFile.txt, то она должна написать два соединённых файла в результирующий.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func writeFile(path, text string) {
	err := os.WriteFile(path, []byte(text), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	switch len(args) {
	case 1:
		str := readFile(args[0])

		fmt.Println(str)
	case 2:
		str1 := readFile(args[0])
		str2 := readFile(args[1])
		result := []string{str1, str2}

		fmt.Print(strings.Join(result, "\n"))
	case 3:
		str1 := readFile(args[0])
		str2 := readFile(args[1])
		result := []string{str1, str2}

		writeFile(args[2], strings.Join(result, "\n"))
	default:
		fmt.Println("Кол-во переданых аргументов не корректно. Требуется от 1 до 3")
	}
}
