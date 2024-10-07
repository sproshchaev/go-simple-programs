package main

import "fmt"

// IsPowerOfTwoRecursive проверяет, является ли N точной степенью двойки
func IsPowerOfTwoRecursive(N int) {
    // Базовый случай: 1 является 2^0
    if N == 1 {
        fmt.Println("YES")
        return
    }
    // Если N меньше или равно 0 или нечетное, возвращаем NO
    if N <= 0 || N%2 != 0 {
        fmt.Println("NO")
        return
    }
    // Рекурсивный вызов: делим N на 2
    IsPowerOfTwoRecursive(N / 2)
}

func main() {
    N := 16 // Пример числа
    IsPowerOfTwoRecursive(N) // Ожидаемый вывод: YES

    N = 18 // Пример числа
    IsPowerOfTwoRecursive(N) // Ожидаемый вывод: NO
}