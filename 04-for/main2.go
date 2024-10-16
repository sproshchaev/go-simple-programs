package main

import (
    "fmt"
)

// напишите программу, которая запрашивает у пользователя число и выводит на экран таблицу умножения от 1 до 10 для этого числа с помощью цикла for.
func main() {
    var number int
    _, err := fmt.Scanf("%d", &number)
    if err != nil {
        return
    }
    for i := 1; i <=10; i++ {
        fmt.Printf("%d * %d = %d\n", number, i, number*i)
    }
}
