package main

import (
	"strings"
	"unicode"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", DeleteVowels("qwerty"))
}

// Напишите на Go тест для функции DeleteVowels(s string) string, которая должна удалять все гласные из строки английского языка (y не считается гласной).
// Используйте table driven testing.
func DeleteVowels(s string) string {
	const vowelsUpperCase = "AEIOU"
	result := "" 
	for i := 0; i < len(s); i++ {
		if strings.IndexRune(vowelsUpperCase, unicode.ToUpper(rune(s[i]))) == -1 {
			result = result + string(s[i]) 
		}
	}
	return result
}
