package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	PrettyArrayOutput(arr)
}

// Напишите функцию PrettyArrayOutput(array [9]string), которая получает массив из девяти строк и печатает:
// "1 я уже сделал: *элемент массива*" для первых 7 элементов
// и "9 не успел сделать: *элемент массива*" для последних двух.
// У строчки должен быть именно такой вид (без кавычек, а вместо "*элемент массива*" сам элемент) и нумерация с единицы.
func PrettyArrayOutput(array [9]string) {
	for i := 0; i < 7; i++ {
		fmt.Println(strconv.Itoa(i+1) + " я уже сделал: " + array[i])
	}
	for i := 7; i < 9; i++ {
		fmt.Println(strconv.Itoa(i+1) + " не успел сделать: " + array[i])
	}
}
