package main

	
import (
    "fmt"
    "time"
)

func main() {
	task := Task{
		summary: "",
		description: "",
		deadline: time.Now(),
		priority: 5, 
	}
	fmt.Printf("%t\n", task.IsOverdue())
	fmt.Printf("%t\n", task.IsTopPriority())
}

// Создайте структуру Task:
// 1. summary - короткий заголовок (тип string)
// 2. description - подробное описание (тип string)
// 3. deadline - дедлайн для задачи (тип time.Time (стандартный пакет time))
// 4. priority - приоритет задачи (тип int)
type Task struct {
	summary string
	description string
	deadline time.Time
	priority int
}

// Для этой структуры реализуйте метод IsOverdue, которая возвращает true, если дедлайн задачи уже прошел и false в ином случае. 
// Подсказка: воспользуйтесь функцией time.Now().
func (t Task) IsOverdue() bool {
    // Не работает time.Now() > t.deadline: оператор > не может быть применен к значениям типа time.Time, потому что это структура, а не примитивный тип данных. Для сравнения двух объектов типа time.Time в Go нужно использовать специальные методы, такие как Before(), After(), или Equal(). 
	if time.Now().After(t.deadline) {
		return true
	} else {
		return false
	}
}

// Ещё запрограммируйте метод IsTopPriority. Функция возвращает true, если приоритет 4 или 5, и false, если меньше.
func (t Task) IsTopPriority() bool {
	if t.priority == 4 || t.priority == 5  {
		return true
	} else {
		return false
	}
}