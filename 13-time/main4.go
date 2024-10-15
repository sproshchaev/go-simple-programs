package main

import (
	"fmt"
	"time"
)

func main() {
	// Пример использования функции
	pastTime := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
	result := TimeAgo(pastTime)
	fmt.Println(result) // Output: 1 month ago (или "X months ago", в зависимости от текущей даты)

	// Пример работы с часами
	pastTime = time.Now().Add(-2 * time.Hour)
	timeAgo := TimeAgo(pastTime)
	fmt.Println(timeAgo) // Output: 2 hours ago
}

// Напишите функцию TimeAgo(pastTime time.Time) string, который вычисляет время, прошедшее с момента заданного значения time.Time,
// и возвращает удобочитаемую строку, указывающую, как давно это было.
// Примечания Пример работы функции:
//
//	pastTime := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
//	result := TimeAgo(pastTime)
//	fmt.Println(result) // Output: 1 month ago
//
// Пример работы с часами:
//
//	pastTime := time.Now().Add(-2 * time.Hour)
//	timeAgo := TimeAgo(pastTime)
//	expected := "2 hours ago"
func TimeAgo(pastTime time.Time) string {
	// Вычисляем разницу между текущим временем и pastTime
	duration := time.Since(pastTime)

	// Определяем единицы времени
	seconds := int(duration.Seconds())
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24
	months := days / 30 // Приблизительное значение
	years := days / 365 // Приблизительное значение

	// Формируем удобочитаемую строку
	switch {
	case years > 0:
		return fmt.Sprintf("%d year%s ago", years, plural(years))
	case months > 0:
		return fmt.Sprintf("%d month%s ago", months, plural(months))
	case days > 0:
		return fmt.Sprintf("%d day%s ago", days, plural(days))
	case hours > 0:
		return fmt.Sprintf("%d hour%s ago", hours, plural(hours))
	case minutes > 0:
		return fmt.Sprintf("%d minute%s ago", minutes, plural(minutes))
	default:
		return fmt.Sprintf("%d second%s ago", seconds, plural(seconds))
	}
}

// plural добавляет 's' к слову, если количество больше 1
func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}
