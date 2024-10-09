package main

import (
	"fmt"
)

func main() {
	fmt.Println(Clean([]int{1, 2, 3, 4, 5, 3}, 3))
}

// Дан неотсортированный слайс целых чисел. Напишите функцию Clean(nums []int, x int) ([]int), которая удаляет из исходного слайса все вхождения x.
// Важно сохранить порядок следования элементов и не использовать дополнительный слайс.
func Clean(nums []int, x int) []int {
	i := 1
	// Эмуляция цикла while с переменной i
	for i != 0 {
		// Проверяем - есть ли элемент?
		if Find(nums, x) != -1 {
			nums = Remove(nums, x)
		} else {
			// Завершаем цикл
			i = 0
		}
	}
	return nums
}

// Поиск позиции элемента в слайсе
func Find(nums []int, x int) int {
	result := -1
	for i := 0; i < len(nums); i++ {
		if x == nums[i] {
			result = i
		}
	}
	return result
}

// Удаление элемента из слайса
func Remove(nums []int, x int) []int {
	findIndex := Find(nums, x)
	if findIndex != -1 {
		                                                  // вариативный оператор ... для соединения срезов
		nums = append(nums[:findIndex], nums[findIndex+1:]...)
	}
	return nums
}