package main

import(
	"fmt"
)

func main() {
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(FindMinMaxInArray(arr))
}

// напишите функцию FindMinMaxInArray(array [10]int) (int, int), которая получает массив 
// из десяти случайных целочисленных значений. Далее функция должна находить 
// максимальное и минимальное значения в массиве и возвращать их.
func FindMinMaxInArray(array [10]int) (int, int) {
	max := array[0] 
	min := array[0]

	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
		if array[i] < min {
			min = array[i]
		}
	}

	return max, min
}