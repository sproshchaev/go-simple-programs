package main

import (
	"fmt"
	"strings"
)

func main() {
	testStrings := []string{
		"a b c c b a",
		"hello",
		"racecar",
		"Madam Im Adam",
		"step on no pets",
	}

	for _, str := range testStrings {
		if IsPalindrome(str) {
			fmt.Printf("\"%s\" является палиндромом.\n", str)
		} else {
			fmt.Printf("\"%s\" не является палиндромом.\n", str)
		}
	}
}

// Напишите функцию IsPalindrome(input string) bool, которая принимает строку и проверяет, является ли она палиндромом.
// Примечания Например, функция IsPalindrome("a b c c b a") должна вернуть true. Вам пригодится пакет strings.
func IsPalindrome(input string) bool {
	// Удаляем пробелы и переводим строку в нижний регистр
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ToLower(input)

	// Инициализируем два указателя: один с начала строки, другой с конца
	left, right := 0, len(input)-1

	for left < right {
		if input[left] != input[right] {
			return false // Если символы не равны, строка не палиндром
		}
		left++
		right--
	}
	return true // Если все символы совпали, строка палиндром
}
