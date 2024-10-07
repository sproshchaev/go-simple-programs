package main

import "fmt"

func main() {
    fmt.Print(SumDigitsRecursive(100))
}

// Напишите функцию SumDigitsRecursive(n int) int, которая рекурсивно вычисляет сумму цифр числа
func SumDigitsRecursive(n int) int {
        // Базовый случай: если число меньше 10, просто возвращаем его
		if n < 10 {
			return n
		}
		return n%10 + SumDigitsRecursive(n/10)
}