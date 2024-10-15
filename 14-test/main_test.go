package main

import (
	"testing"
)

// TestSum - тест для функции Sum.
// Запуск тестов go test
func TestSum(t *testing.T) {
	// Тест с положительными числами
	result := Sum(1, 1)
	if result != 2 {
		t.Errorf("Sum(1, 1) = %d; want 2", result)
	}

	// Тест с нулями
	result = Sum(0, 0)
	if result != 0 {
		t.Errorf("Sum(0, 0) = %d; want 0", result)
	}

	// Тест с отрицательными числами
	result = Sum(-1, -1)
	if result != -2 {
		t.Errorf("Sum(-1, -1) = %d; want -2", result)
	}

	// Тест с положительным и отрицательным числом
	result = Sum(1, -1)
	if result != 0 {
		t.Errorf("Sum(1, -1) = %d; want 0", result)
	}

	// Тест с большими числами
	result = Sum(100, 200)
	if result != 300 {
		t.Errorf("Sum(100, 200) = %d; want 300", result)
	}
}