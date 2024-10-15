package main

import (
	"errors"
	"fmt"
)

func main() {
	num := 5
	result, err := Factorial(num)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Факториал %d = %d\n", num, result)
	}
}

// Напишите функцию Factorial(n int) (int, error) для нахождения факториала без ошибок.
// Функция получает на вход целое число и выводит его факториал. Если пользователь ввёл отрицательное число,
// программа должна вернуть сообщение об ошибке (factorial is not defined for negative numbers).
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	result := 1
	for i := 2; i <= n; i++ { // Факториал 0 и 1 равен 1
		result *= i
	}
	return result, nil
}
