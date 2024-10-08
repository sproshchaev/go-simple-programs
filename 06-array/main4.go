package main

import (
	"fmt"
)

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(SumOfArray(arr))
}

// Напишите функцию SumOfArray(array [6]int) int которая получает массив из шести целочисленных элементов и возвращает сумму всех элементов массива.
func SumOfArray(array [6]int) int {
	result := 0
	for i := 0; i < len(array); i++ {
		result = result + array[i]
	}
	return result
}
