package main

import (
        "fmt"
)

// Напишите программу, которая запрашивает у пользователя два числа и выводит на экран сообщение: 
// Первое число больше второго, Второе число больше первого и Числа равны по результатам сравнения.
func main() {
    var number1, number2 int
    _, err := fmt.Scanf("%d %d", &number1, &number2)

    if err != nil {
        fmt.Println("Ошибка", err)
        return
    }
     if number1 > number2 {
        fmt.Println("Первое число больше второго")
    } else if number1 < number2 {
        fmt.Println("Второе число больше первого")
    } else {
        fmt.Println("Числа равны")
    } 
}
