package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%s\n", ConcatStringsAndInt("Hello", "World", 2024))
}

// Напишите функцию ConcatStringsAndInt(str1, str2 string, num int) string, которая принимает две строки и одно целое число, а затем выполняет конкатенацию строк и числа в одну строку.
func ConcatStringsAndInt(str1, str2 string, num int) string {
	return str1 + str2 + strconv.Itoa(num)
}
