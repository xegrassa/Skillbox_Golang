/*
Написать программу для нахождения подстроки в кириллической подстроке. Программа должна запускаться с помощью команды:

	go run main.go --str "строка для поиска" --substr "поиска"

Для реализации такой работы с флагами воспользуйтесь пакетом flags, а для поиска подстроки в строке вам понадобятся руны.

Что нужно сделать

	Спроектировать алгоритм поиска подстроки.
	Определить строку и подстроку, используя флаги.
	Написать алгоритм реализацию для работы со строками UTF-8 (для этого необходимо воспользоваться рунами).

Что оценивается

	Алгоритм может работать с различными символами (кириллица, китайские иероглифы).
	Использованы руны.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"unicode/utf8"
)

func isContain(sentence, substring string) bool {
	firstSubRune, _ := utf8.DecodeRuneInString(substring)
	runeCount := utf8.RuneCountInString(substring)

	log.Printf("Строка: %v | Подстрока: %v\n", sentence, substring)

	for sentenceIndex, sentenceRune := range sentence {

		isContain := true
		if firstSubRune == sentenceRune {

			sentenceSlice := sentence[sentenceIndex:]
			var rSize int
			for i := 0; i < runeCount; i++ {
				srune, _ := utf8.DecodeRuneInString(sentenceSlice[rSize:])
				subrune, subsize := utf8.DecodeRuneInString(substring[rSize:])
				log.Printf("%v compare: Sentence=%c Sub=%c\n", i+1, srune, subrune)

				if srune != subrune {
					isContain = false
					break
				}
				rSize += subsize
			}

			if isContain {
				return true
			}
		}
	}
	return false
}

func main() {
	var (
		sentence  string
		substring string
	)
	flag.StringVar(&sentence, "str", "", "Строка в которой искать подстроку")
	flag.StringVar(&substring, "substr", "", "Подстрока")
	flag.Parse()

	fmt.Println(isContain(sentence, substring))
}
