package main

import (
	"testing"
)

// запустить go test -cover
type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"}, // Тест на отрицательное число
	{0, "zero"},      // Тест на ноль
	{5, "short"},     // Тест на короткое число
	{50, "long"},     // Тест на длинное число
	{100, "very long"}, // Тест на очень длинное число
	{1000, "very long"}, // Еще один тест на очень длинное число
}

func TestLength(t *testing.T) {
	for i, test := range tests {
		size := Length(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}