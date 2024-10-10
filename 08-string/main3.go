package main

import(
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", isLatin("qwerty"))

}

// Напишите функцию isLatin(input string) bool, которая принимает строку и выводит true, если все символы в строке латинские, false, если нет.
// Подсказка: советуем использовать стандартную библиотеку unicode
func isLatin(input string) bool {
	for _, r := range input {
		if !unicode.Is(unicode.Latin, r) {
			return false
		}
	}
	return true
}
