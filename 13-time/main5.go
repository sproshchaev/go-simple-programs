package main

import (
	"fmt"
	"time"
)

func main() {
	// Пример использования функции
	start := time.Date(2023, 10, 20, 0, 0, 0, 0, time.UTC) // Пятница
	nextWorkday := NextWorkday(start)
	fmt.Println("Next workday:", nextWorkday) // Output: 2023-10-23 00:00:00 +0000 UTC (понедельник)

	start = time.Date(2023, 10, 21, 0, 0, 0, 0, time.UTC) // Суббота
	nextWorkday = NextWorkday(start)
	fmt.Println("Next workday:", nextWorkday) // Output: 2023-10-23 00:00:00 +0000 UTC (понедельник)

	start = time.Date(2023, 10, 22, 0, 0, 0, 0, time.UTC) // Воскресенье
	nextWorkday = NextWorkday(start)
	fmt.Println("Next workday:", nextWorkday) // Output: 2023-10-23 00:00:00 +0000 UTC (понедельник)
}

// Напишите функцию NextWorkday(start time.Time) time.Time, которая вычисляет дату следующего рабочего дня (исключая выходные). 
// Учитывая дату начала, функция должна возвращать дату следующего рабочего дня.
// Примечания Рабочая неделя начинается с понедельника :)
func NextWorkday(start time.Time) time.Time {
	nextDay := start.AddDate(0, 0, 1) // Переход к следующему дню
	// Проверяем, является ли следующий день выходным (суббота или воскресенье)
	for nextDay.Weekday() == time.Saturday || nextDay.Weekday() == time.Sunday {
		nextDay = nextDay.AddDate(0, 0, 1) // Переход к следующему дню
	}
	return nextDay
}