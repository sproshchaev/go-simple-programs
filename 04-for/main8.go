package main

import (
	"fmt"
)

// Напишите программу, которая запрашивает у пользователя число и выводит на экран 10 чисел Фибоначчи, начиная с введённого пользователем, если оно является таким числом, или с ближайшего большего.
// Числа Фибоначчи — последовательность чисел, в которой каждое следующее число равно сумме двух предыдущих: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34...
// Формат ввода Целое число. Формат вывода Целые числа. Каждое число – с новой строки.

// Функция для генерации 10 чисел Фибоначчи, начиная с заданного числа
func fibonacci(start int) []int {
	fibNumbers := make([]int, 0, 10)

	// Начальные значения последовательности Фибоначчи
	a, b := 0, 1
	for len(fibNumbers) < 10 {
		if b >= start {
			fibNumbers = append(fibNumbers, b)
		}
		a, b = b, a+b
	}

	return fibNumbers
}

func main() {
	var input int

	// Запрос числа у пользователя
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	// Получаем числа Фибоначчи
	fibNumbers := fibonacci(input)

	// Выводим результаты
	for _, num := range fibNumbers {
		fmt.Println(num)
	}
}
