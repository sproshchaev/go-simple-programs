package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	


}

// Напишите функцию SumTwoIntegers(a, b string) (int, error), которая получает две строки и,
// если это целые числа, возвращает их сумму. Если пользователь ввёл данные не целого типа,
// функция должна вернуть сообщение об ошибке (invalid input, please provide two integers).
func SumTwoIntegers(a, b string) (int, error) {
	
	aInt, aErr := strconv.Atoi(a)
	bInt, bErr := strconv.Atoi(b)
	
	if aErr != nil || bErr != nil  {
		return 0, errors.New("invalid input, please provide two integers")
	} else {
		return aInt + bInt, nil
	}

}