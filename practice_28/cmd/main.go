/*
Цель задания

	Научиться работать с композитными типами данных: структурами и картами

Что нужно сделать

	Напишите программу, которая считывает ввод с stdin, создаёт структуру student и записывает указатель на структуру в хранилище map[studentName] *Student.

	Программа должна получать строки в бесконечном цикле, создать структуру Student через функцию newStudent,
	далее сохранить указатель на эту структуру в map, а после получения EOF (ctrl + d) вывести на экран имена всех студентов из хранилища.
	Также необходимо реализовать методы put, get.

Пример:
Строки

	Вася 24 1
	Семен 32 2
	EOF

Output

	Студенты из хранилища:
	Вася 24 1
	Семен 32 2

Критерии оценки
Зачёт:

	при получении одной строки (например, «имяСтудента 24 1») программа создаёт студента и сохраняет его, далее ожидает следующую строку или сигнал EOF (Сtrl + Z);
	при получении сигнала EOF программа должна вывести имена всех студентов из map.
*/
package main

import (
	"fmt"
	"io"
	"skillbox/practice_28/pkg/storage"
	"skillbox/practice_28/pkg/student"
	"strings"
)

func main() {
	var (
		name     string
		age      int
		grade    int
		sPointer *student.Student
	)
	st := storage.NewStorage()

	fmt.Println("Введите имя, возраст, грейд:")
Loop:
	for {
		_, err := fmt.Scanf("%s %d %d", &name, &age, &grade)

		switch err {
		case io.EOF:
			break Loop
		case nil:
			sPointer = student.NewStudent(name, age, grade)
			st.Put(sPointer)
		default:
			fmt.Println("Введены некорректные данные. Введите еще раз. Для выхода нажмите CTRL+D")
		}
	}

	fmt.Println("Студенты из хранилища:")
	str := strings.Join(st.Get(), "\n")
	fmt.Println(str)
}
