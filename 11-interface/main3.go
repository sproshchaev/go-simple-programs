package main

import(
	"fmt"
)

func main() {
	// errorLog := &Log{logLevel: Error}
	errorLog := &Log{Level: Error}
	errorLog.Log("This is an error message")
}

// Создайте интерфейс Logger с методом Log, который будет записывать сообщение в журнал. 
type Logger interface {
	Log(s string)
}

// Создайте пользовательский тип LogLevel типа string, сделайте константные значения типа LogLevel со значениями Error и Info. 
type LogLevel string

// Определяем константы для уровней логирования
const (
	Error LogLevel = "Error"
	Info  LogLevel = "Info"
)

// Создайте структуру Log с полем LogLevel. 
type Log struct {
	//logLevel LogLevel
	Level LogLevel
}

// Реализуйте метод Log c параметром типа string (текст ошибки), который в зависимости от текущего LogLevel выводит сообщение "ERROR: *текст ошибки*" или "INFO: *текст ошибки*".
func (l Log) Log(s string) {
	if l.Level == "Info" {
		fmt.Printf("INFO: %s\n", s)
	} else {
		fmt.Printf("ERROR: %s\n", s)
	}
}