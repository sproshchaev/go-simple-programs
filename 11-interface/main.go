package main

import(
	"fmt"
)

fuct main() {

}

// Создайте интерфейс Shape с методом Area, который будет возвращать площадь фигуры. 
type Shape interface {
    Area() float64
}

// Создайте структуры Circle и Rectangle, которые будут реализовывать этот интерфейс и рассчитывать площадь этих фигур.
// Примечания: Вам может пригодиться пакет math.
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return (3.14159 * c.radius * c.radius)  
}

type Rectangle struct {
	height float64
	width float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}
