package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now()
	format := "2006-01-02 15:04:05" // Год-Месяц-День Час:Минута:Секунда
	formattedTime := FormatTimeToString(timestamp, format)
	fmt.Println(formattedTime)
}

// Создайте функцию FormatTimeToString(timestamp time.Time, format string) string,
// которая форматирует заданное значение time.Time как строку в определенном формате.
// Функция должна принимать значение time.Time и строку формата и возвращать форматированную строку.
// Примечение: Пример работы функции:
// timestamp := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
// format := "2006-01-02 15:04:05"
// result := FormatTimeToString(timestamp, format)
// fmt.Println(result) // Output: 2023-10-23 02:41:49

func FormatTimeToString(timestamp time.Time, format string) string {
	return timestamp.Format(format)
}
