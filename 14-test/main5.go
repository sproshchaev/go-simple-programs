package main

import (
	//"fmt"
	"errors"
	"unicode/utf8"
)

//func main() {
//}

// Функция GetUTFLength(input []byte) (int, error) возвращает длину строки UTF8 и ошибку ErrInvalidUTF8 (в случае возникновения)
func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, errors.New("invalid utf8")
	} else {

		return utf8.RuneCount(input), nil
	}
	return 0, nil
}
