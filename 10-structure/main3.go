package main

import (
	"fmt"
)

func main() {
	student := Student{
		name:            "Ivan",
		solvedProblems:  100,
		scoreForOneTask: 1,
		passingScore:    100,
	}
	fmt.Printf("%t\n", student.IsExcellentStudent())

}

// Создайте структуру Student с полями name (строка),
// solvedProblems — количество решённых задач (целое число),
// scoreForOneTask— количество баллов за одну задачу; будем считать, что все задачи дают одинаковые баллы (число с плавающей точкой)
// и passingScore — проходной балл в следующий модуль (число с плавающей точкой).
type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

// Создайте метод IsExcellentStudent для этой структуры, который возвращает true,
// если ученик проходит по баллам в следующий модуль и false в ином случае.
func (s Student) IsExcellentStudent() bool {
	if float64(s.solvedProblems)*s.scoreForOneTask >= s.passingScore {
		return true
	} else {
		return false
	}

}
