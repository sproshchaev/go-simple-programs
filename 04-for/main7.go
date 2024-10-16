package main

import (
    "fmt"
)

// Напишите программу она должна запрашивать у пользователя число и выводить на экран сумму всех чисел от 1 до этого числа, кроме чисел, 
// которые делятся на 3 или на 5 с помощью цикла for.
func main() {
    var n int
    _, err := fmt.Scanf("%d", &n)
    if err != nil || n < 1 {
        return
    }
    sum := 0
    for i:= 1; i <= n; i++ {
        if i%3 != 0 && i%5 != 0 {
            sum += i
        }
    }
    fmt.Println(sum)
}