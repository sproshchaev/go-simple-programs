package main

import(
	"fmt"
)

func main() {
	fmt.Printf("Результат %s\n", ConcatenateStrings("Hello,", "World!"))
}

// Напишите функцию ConcatenateStrings(str1, str2 string) string, которая принимает две строки и склеивает их вместе, разделяя пробелом.
func ConcatenateStrings(str1, str2 string) string {
	return str1 + " " + str2
}