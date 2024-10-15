package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	num := 5
	result, err := IntToBinary(num)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Двоичное представление %d: %s\n", num, result)
	}
}

// Напишите функцию IntToBinary(num int) (string, error), которая принимает на вход целое число и возвращает его двоичное представление.
// Если пользователь ввёл отрицательное число, программа должна возвращать сообщение об ошибке (negative numbers are not allowed).
// Примечания: Обратите внимание, что функция возвращает строку, а не число. В случае ошибки верните пустую строку. Много стандартных операций уже реализовано во внешних пакетах.
func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", errors.New("negative numbers are not allowed")
	}

	return strconv.FormatInt(int64(num), 2), nil
}
