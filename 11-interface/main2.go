package main

import("fmt")

func main() {

	// Создание объектов Dog и Cat
	dog := Dog{name: "Бобик"}
	cat := Cat{name: "Мурка"}

	// Тестирование интерфейса и вызов метода MakeSound
	var animal Animal

	// Присваиваем Dog переменной интерфейса Animal и вызываем MakeSound
	animal = dog
	animal.MakeSound()

	// Присваиваем Cat переменной интерфейса Animal и вызываем MakeSound
	animal = cat
	animal.MakeSound()

}

// Создайте интерфейс Animal с методом MakeSound, который будет выводить звук, издаваемый животным. 
type Animal interface {
	MakeSound()
}

// Создайте структуры Dog и Cat, которые будут реализовывать этот интерфейс и издавать соответствующие звуки (выводить на экран) – "Гав!" и "Мяу!".
type Dog struct {
	name string
}

func (d Dog) MakeSound() {
	fmt.Printf("%s\n", "Гав!")
}

type Cat struct {
	name string
}

func (c Cat) MakeSound() {
	fmt.Printf("%s\n", "Мяу!")
}
