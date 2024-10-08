package main

import(
	"fmt"
) 

// Напишите функцию ThirdElementInArray(array [6]int) int, которая получает массив из шести целочисленных элементов (пусть это будут номера игроков на форме) 
// и возвращает третий по счёту элемент массива.
func main() {
	input := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(ThirdElementInArray(input))

}

func ThirdElementInArray(array [6]int) int {
	return array[2]
}