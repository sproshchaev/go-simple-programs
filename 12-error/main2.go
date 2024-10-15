package main

import (
	"fmt"
	"errors"
)

func main() {
	a := 100
	b := 20
	result, err := DivideIntegers(a, b)
    if err != nil {
        fmt.Printf("Ошибка: %s\n", err) // Вывод ошибки
        return
    } else {
		fmt.Printf("Результат деления %d на %d: %.2f\n", a, b, result) // Вывод результата
	}
}

// Напишите функцию DivideIntegers(a, b int) (float64, error), которая возвращает результат деления первого числа на второе.
// Если второе число равно нулю, функция должна возвращать в качестве ответа 0 и сообщение об ошибке (division by zero is not allowed).
// Если второе число не равно нулю – верните в качестве ошибки nil.
func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		// возвращать в качестве ответа 0 и сообщение об ошибке (division by zero is not allowed)
		return 0, errors.New("division by zero is not allowed")
	} else {
		// Если второе число не равно нулю – верните в качестве ошибки nil.
		return float64(a / b), nil
	}

}
