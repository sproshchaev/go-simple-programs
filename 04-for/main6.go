package main

import (
    "fmt"
)

// Напишите программу, которая запрашивает у пользователя число и выводит на экран все числа от 1 до этого числа, которые делятся на 3 с помощью цикла for.
func main() {
    var n int
    _, err := fmt.Scanf("%d", &n)
    if err != nil || n < 1{
        return
    }
    for i := 1; i <= n; i++ {
        if i%3 == 0 {
            fmt.Println(i)
        }
    }
}
