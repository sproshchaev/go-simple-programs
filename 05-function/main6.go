package main

import (
	"fmt"
	"math"
)

func main() {
	k1, b1 := 2.0, 1.0
	k2, b2 := -1.0, 4.0

	x, y := FindIntersection(k1, b1, k2, b2)

	if math.IsNaN(x) && math.IsNaN(y) {
		fmt.Println("Прямые параллельны, пересечения нет.")
	} else {
		fmt.Printf("Точка пересечения: (%.2f, %.2f)\n", x, y)
	}
}

// Напишите функцию FindIntersection(k1, b1, k2, b2 float64) (float64, float64), 
// которая будет принимать на вход четыре вещественных числа k1, b1, k2 и b2 — они представляют коэффициенты уравнений y = kx + b 
// для двух разных астрономических объектов. Ваша задача — найти точку, в которой эти астрономические траектории пересекаются. 
// Если траектории параллельны, программа должна вернуть math.NaN(), math.NaN().
func FindIntersection(k1, b1, k2, b2 float64) (float64, float64) {
	// Проверка на параллельность
	if k1 == k2 {
		return math.NaN(), math.NaN() // Параллельные прямые не пересекаются
	}

	// Вычисление координат точки пересечения
	x := (b2 - b1) / (k1 - k2)
	y := k1*x + b1

	return x, y
}
