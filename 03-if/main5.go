package main

import (
    "fmt"
)

// Напишите программу, которая запрашивает у пользователя три числа и выводит на экран сообщение: 
// Все числа равны, Два числа равны и Все числа разные по результату их сравнения. Если ни одно 
// из условий не выполняется, программа должна выводить сообщение: Некорректный ввод.
func main() {
    var number1, number2, number3 int
    n, err := fmt.Scanf("%d %d %d", &number1, &number2, &number3)
    if err != nil || n != 3 {
        fmt.Println("Некорректный ввод")
        return
    }
    if number1 == number2 && number2 == number3 {
        fmt.Println("Все числа равны")
    } else if number1 == number2 || number1 == number3 || number2 == number3 {
        fmt.Println("Два числа равны")
    } else {
        fmt.Println("Все числа разные")
    }
}