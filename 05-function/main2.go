package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Multiply(10, 10))
	PrintNumbersAscending(10)
}

// Add(a, b float64) float64 Возвращает a + b
func Add(a, b float64) float64 {
    return a + b
}

// Multiply(a, b float64) float64 Возвращает a * b
func Multiply(a, b float64) float64 {
    return a * b
}

// PrintNumbersAscending(n int) Печатает возрастающую последовательность от 1 до (включая) n через пробел
func PrintNumbersAscending(n int) {
	for i := 1; i<=n; i++ {
		fmt.Print(strconv.Itoa(i) + " ")
	}
}
