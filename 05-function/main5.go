package main

import (
	"fmt"
)

func main() {
	var n int

	// Запрос числа у пользователя
	fmt.Print("Введите целое число n: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	// Вычисляем сумму последовательности
	result := CalculateSeriesSum(n)

	// Выводим результат
	fmt.Printf("Сумма последовательности для n = %d: %.6f\n", n, result)
}

// Написать функцию CalculateSeriesSum(n int) float64, которая получает на вход число n и вычисляет сумму указанной последовательности. 
// Затем программа должна вернуть полученное значение.
func CalculateSeriesSum(n int) float64 {
	var sum float64 = 0.0
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i)
	}
	return sum
}