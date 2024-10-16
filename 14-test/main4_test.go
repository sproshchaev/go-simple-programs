package main

import (
	"testing"
)

// Напишите на Go тест для функции DeleteVowels(s string) string, которая должна удалять все гласные из строки английского языка (y не считается гласной).
// Используйте table driven testing.
// Тестовая функция должна начинаться с Test
func TestDeleteVowels(t *testing.T) {
	tests := []struct{
		input string
		expected string
	}{
		{"abc", "bc"},
		{"qwerty", "qwrty"},
		{"igls", "gls"},
		{"ot", "t"},
		{"uk", "k"},
	} 

	// Цикл по тестовым случаям
	for _, test := range tests {
		// Вызов функции, которую тестируем:
		result := DeleteVowels(test.input)
		// Анализ результата 
		if result != test.expected {
			// Регистрация ошибки в тестовом фреймфорке через t.Errorf
			t.Errorf("Test fail! DeleteVowels(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}

}