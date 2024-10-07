package main

import (
	"fmt"
	"math"
)

func main() {
	SqRoots()
}

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	// Вычисление дискриминанта
	d := b*b - 4*a*c

	// Проверка дискриминанта
	if d < 0 {
		fmt.Println("0 0") // Нет корней
		return
	}

	root1 := (-b + math.Sqrt(d)) / (2 * a)
	root2 := (-b - math.Sqrt(d)) / (2 * a)

	// Вывод корней
	if d == 0 {
		// Один корень
		fmt.Printf("%.6f\n", root1)
	} else {
		// Два корня, выводим в порядке возрастания
		if root1 > root2 {
			root1, root2 = root2, root1
		}
		fmt.Printf("%.6f %.6f\n", root1, root2)
	}
}
