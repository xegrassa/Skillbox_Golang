package storage

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func newStudent(name string, age, grade int) *Student {
	return &Student{name, age, grade}
}

type Storage map[string]*Student

func (st Storage) Get() []string {
	result := make([]string, 0, 5)
	var str string
	for _, s := range st {
		str = fmt.Sprintf("%v %v %v", s.name, s.age, s.grade)
		result = append(result, str)
	}
	return result
}

func (st Storage) Put(name string, age, grade int) {
	st[name] = newStudent(name, age, grade)
}
