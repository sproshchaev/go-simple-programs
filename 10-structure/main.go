package main

import(
	"fmt"
)

func main() {

	// Создаем экземпляр структуры
    person := Person{name: "Alice", age: 30, address: "NY"}

    // Вызываем метод Print
    person.Print()
}

// Создайте структуру Person с полями name, age и address. 
type Person struct {
	name string 
	age int
	address string
}

// Создайте метод Print для этой структуры, который будет выводить информацию о человеке на экран в формате:
// Name: Гоша
// Age: 21
// Address: Ясногорск
func (p Person) Print() {
	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Age: %d\n", p.age)
	fmt.Printf("Address: %s\n", p.address)
} 
