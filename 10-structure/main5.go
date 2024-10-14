package main

import(
	"fmt"
	"time"
)

func main() {

	// Создание заметок
	note1 := Note{title: "Заметка 1", text: "Текст первой заметки"}
	note2 := Note{title: "Заметка 2", text: "Текст второй заметки"}

	// Создание задач
	task1 := Task{
		summary:     "Задача 1",
		description: "Описание задачи 1",
		deadline:    time.Now().Add(48 * time.Hour), // Дедлайн через 2 дня
		priority:    5,
	}
	task2 := Task{
		summary:     "Задача 2",
		description: "Описание задачи 2",
		deadline:    time.Now().Add(-24 * time.Hour), // Просроченная задача (дедлайн 1 день назад)
		priority:    3,
	}
	task3 := Task{
		summary:     "Задача 3",
		description: "Описание задачи 3",
		deadline:    time.Now().Add(24 * time.Hour), // Дедлайн через 1 день
		priority:    4,
	}

	// Создание списка дел
	toDoList := ToDoList{
		name:  "Список дел на сегодня",
		tasks: []Task{task1, task2, task3},
		notes: []Note{note1, note2},
	}

	// Вызов методов
	fmt.Printf("Название списка: %s\n", toDoList.name)
	fmt.Printf("Количество задач: %d\n", toDoList.TasksCount())
	fmt.Printf("Количество заметок: %d\n", toDoList.NotesCount())
	fmt.Printf("Количество приоритетных задач: %d\n", toDoList.CountTopPrioritiesTasks())
	fmt.Printf("Количество просроченных задач: %d\n", toDoList.CountOverdueTasks())
}

// Напишите структура Note (сущность заметок, у которых в отличие от задач нет чётких дедлайнов и приоритета):
// 1. title - заголовок (тип string)
// 2. text - текст заметок (тип string)
type Note struct {
	title string
	text string
}

// Создайте структуру ToDoList с такими полями:
// 1. name - название списка (тип string)
// 2. tasks - список дел на сегодня (тип слайс структур Task (из предыдущего задания))
// 3. notes - список дополнительных заметок (тип слайс структур Note)
type ToDoList struct {
	name string
	tasks []Task
	notes []Note
}

// Для этой структуры реализуйте методы TasksCount 
// и NotesCount, которые возвращают общее количество задач и заметок соответственно.
func (t ToDoList) TasksCount() int {
	return len(t.tasks)
}

func (t ToDoList) NotesCount() int {
	return len(t.notes)
}

// Также реализуйте метод CountTopPrioritiesTasks, который возвращает количество приоритетных задач. 
// А также метод CountOverdueTasks, который возвращает количество просроченных задач.
func (t ToDoList) CountTopPrioritiesTasks() int {
	result := 0
	for _, value := range t.tasks {
		if value.IsTopPriority() {
			result++ 
		}
	}
    return result
}

func (t ToDoList) CountOverdueTasks() int {
	result := 0
	for _, value := range t.tasks {
		if value.IsOverdue() {
			result++
		}
	}
    return result
}
// Сама структура Task и все её методы из предыдущего задания также должны быть реализованы в этом.
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
