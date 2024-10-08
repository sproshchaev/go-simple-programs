package main

import (
	"fmt"
)

// Напишите функцию FiveSteps(array [5]int) [5]int, которая принимает массив из пяти целочисленных элементов.
// Программа должна возвращать значения элементов массива в обратном порядке.
func main() {
	var myArr [5]int
	myArr = [5]int{0, 1, 2, 3, 4}
	result := FiveSteps(myArr)
	fmt.Println(result)
	for i := 0; i < len(result); i++ {
		fmt.Printf("%d ", result[i])
	}
}

func FiveSteps(array [5]int) [5]int {
	var arr [5]int
	for i := len(array) - 1; i >= 0; i-- {
		arr[len(array)-1-i] = array[i]
	}
	return arr
}
