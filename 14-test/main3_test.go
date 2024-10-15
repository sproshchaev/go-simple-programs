package main

import (
	"testing"
)

// TestMultiply - тест для функции Multiply.
// go test                  - проходят
// go test -v main3_test.go - не проходят
func TestMultiply(t *testing.T) {
	// Тест с положительными числами
	result := Multiply(2, 3)
	if result != 6 {
		t.Errorf("Multiply(2, 3) = %d; want 6", result)
	}

	// Тест с отрицательными числами
	result = Multiply(-2, -3)
	if result != 6 {
		t.Errorf("Multiply(-2, -3) = %d; want 6", result)
	}

	// Тест с положительным и отрицательным числом
	result = Multiply(2, -3)
	if result != -6 {
		t.Errorf("Multiply(2, -3) = %d; want -6", result)
	}

	// Тест с нулем
	result = Multiply(0, 5)
	if result != 0 {
		t.Errorf("Multiply(0, 5) = %d; want 0", result)
	}

	result = Multiply(5, 0)
	if result != 0 {
		t.Errorf("Multiply(5, 0) = %d; want 0", result)
	}
}