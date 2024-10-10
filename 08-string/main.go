package main

import(
	"fmt"
)

func main() {
	fmt.Println(StringLength("Hello!"))
}

// Напишите функцию StringLength(input string) int, которая принимает строку и возвращает её длину.
func StringLength(input string) int {
	return len(input)
}