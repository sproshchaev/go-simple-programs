package main

import (
	"fmt"
)

// Напишите функцию FiveSteps(array [5]int) [5]int, которая принимает массив из пяти целочисленных элементов.
// Программа должна возвращать значения элементов массива в обратном порядке.
func main() {
	fmt.Println("Hello!")
	var myArr [5]int
	myArr = [5]int{0, 1, 2, 3, 4}
	fmt.Println(FiveSteps(myArr))
}

func FiveSteps(array [5]int) [5]int {
	var arr [len(array)]int
	for i := len(array) - 1; i >= 0; i-- {
		arr[len(array)-1-i] = array[i]
	}
	return arr
}
