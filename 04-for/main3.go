package main

import(
    "fmt"
)

// Напишите программу, которая запрашивает у пользователя число и выводит на экран сумму всех чисел от 1 до этого числа с помощью цикла for
func main() {
    var number int
    _, err := fmt.Scanf("%d", &number)
    if err != nil || number < 1 {
        return
    }
    sum := 0
    for i := 1; i <= number; i++ {
        sum += i
    }
    fmt.Println(sum)
}