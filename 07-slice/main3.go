package main

import (
	"fmt"
)

func main() {

	// Создаем слайс с длиной 3 и емкостью 10
	nums := make([]int, 3, 10)

	// Заполняем первые три элемента
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3

	// До
	fmt.Printf("ДО: Емкость: %d, Длина: %d\n", cap(nums), len(nums))
	fmt.Println(nums)

	// Новый слайс  
	nums = SliceCopy(nums)

	// После
	fmt.Printf("ПОСЛЕ: Емкость: %d, Длина: %d\n", cap(nums), len(nums))
	fmt.Println(nums)

}

// Дан слайс целых чисел nums. Этот слайс имеет емкость больше его длины.
// Создайте функцию SliceCopy(nums []int) []int, которая вернёт новый слайс длиной и ёмкостью, равной длине nums.
// Скопируйте в него значения из исходного слайса
func SliceCopy(nums []int) []int {

	// Длина слайса
	lenVar := len(nums)

    // Создаем новый слайс с длиной lenVar и емкостью capVar
    newNums := make([]int, lenVar, lenVar)

	// Копируем элементы из nums (source) в newNums (destination)
    copy(newNums, nums)

	return newNums
}
