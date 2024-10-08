package main

import (
	"errors"
	"fmt"
)

func main() {

	input := []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12}
	result, err := UnderLimit(input, 3, 5)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}

}

// Дан неотсортированный слайс целых чисел. Напишите функцию UnderLimit(nums []int, limit int, n int) ([]int, error),
// которая будет возвращать первые n (либо меньше, если остальные не подходят) элементов, которые меньше limit.
// В случае ошибки функция должна вернуть nil и описание ошибки.
func UnderLimit(nums []int, limit int, n int) ([]int, error) {

	// Проверка на nil
	if nums == nil {
		return nil, errors.New("входной слайс равен nil")
	}

	// Проверка на отрицательное значение n
	if n < 0 {
		return nil, errors.New("n не может быть отрицательным")
	}

	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] < limit {
			result = append(result, nums[i])
		}
	}

	if len(result) <= n {
		return result, nil
	} else {
		return nil, errors.New("Ошибка")
	}

}
