package main

import (
    "fmt"
)

// Напишите программу, которая запрашивает у пользователя два числа и выводит на экран сообщение: 
// Оба числа положительные, Оба числа отрицательные, Одно число положительное, а другое отрицательное 
// по результатам их проверки. Выведите на экран сообщение: Одно из чисел равно нулю, если это так.
func main() {
    var number1, number2 int
    _, err := fmt.Scanf("%d %d", &number1, &number2)
    if err != nil {
        return
    }
    if number1 == 0 || number2 == 0 {
        fmt.Println("Одно из чисел равно нулю")
    } else if number1 > 0 && number2 > 0 {
        fmt.Println("Оба числа положительные")
    } else if number1 < 0 && number2 < 0 {
        fmt.Println("Оба числа отрицательные")
    } else {
        fmt.Println("Одно число положительное, а другое отрицательное")
    }
}
