/*
Задание 2. Поиск символов в нескольких строках
Что нужно сделать
Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune,
а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово
в предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:

func parseTest(sentences []string, chars []rune)

Советы и рекомендации
Не забудьте проверить, что вы получили больше чем 0 аргументов.
Подход не важен: можно переписать сортировку пузырьком или отсортировать, а потом перевернуть.
Пример входных данных
sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}

chars := [5]rune{'H','E','L','П','М'}

# Пример вывода результата в первом элементе массива

'H' position 0

'E' position 1

'L' position 9
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func getCharIndexInLastWord(sentence2 string, char rune) int {
	sentence := sentence2
	for len(sentence) > 0 {
		r, size := utf8.DecodeLastRune([]byte(sentence))
		if r == char {
			return len(sentence) - size
		}
		sentence = sentence[:len(sentence)-size]
	}
	return -1
}

func parseTest(sentences []string, chars []rune) [][]int {
	sentencesCount := len(sentences)
	charsCount := len(chars)

	if sentencesCount == 0 {
		panic("Не переданы строки !")
	}
	if charsCount == 0 {
		panic("Не переданы Руны !")
	}

	mas := make([][]int, sentencesCount, sentencesCount)
	for i := range mas {
		mas[i] = make([]int, charsCount, charsCount)
	}

	for i, sentence := range sentences {
		for j, char := range chars {
			mas[i][j] = getCharIndexInLastWord(sentence, char)
		}
	}
	return mas
}

func main() {
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'H', 'e', 'l', 'П', 'р'}

	mas := parseTest(sentences, chars)

	fmt.Println(mas)

	for i, _ := range mas {
		fmt.Printf("В строке: %v\n", sentences[i])
		for j, index := range mas[i] {
			fmt.Printf("%c, position %v\n", chars[j], index)
		}
		fmt.Println()
	}
}
