package main

import(
	"fmt"
)

func main() {
	employee := Employee{
		name: "Arseniy",
		position: "backend developer",
		salary: 1000,
		bonus: 100,
	}
	employee.CalculateTotalSalary()
}

// Создайте структуру Employee с полями name (string), position (string), salary (float64) и bonus (float64)
type Employee struct {
	name string
	position string
	salary float64
	bonus float64
}
// Создайте метод CalculateTotalSalary для этой структуры, который будет выводить общую зарплату работника (salary + bonus) по следующему формату:
// Employee: Arseniy
// Position: backend developer
// Total Salary: 1000.02 - Обратите внимание, что Total Salary нужно округлять до 2 знаков после запятой.
func (e Employee) CalculateTotalSalary() {
	fmt.Printf("Employee: %s\n", e.name)
	fmt.Printf("Position: %s\n", e.position)
	fmt.Printf("Total Salary: %s\n", fmt.Sprintf("%.2f", e.salary + e.bonus))
}
