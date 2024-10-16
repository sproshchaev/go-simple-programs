package main

import("testing")

func TestGetUTFLength(t *testing.T) {
	tests:= []struct {
		input    []byte
		expected int
		hasError bool
	}{
		{[]byte("hello"), 5, false},               // обычный случай
		{[]byte{0xff}, 0, true},                   // некорректный UTF-8
	}
	for _, test := range tests {
		result, err := GetUTFLength(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("GetUTFLength(%q) unexpected error: %v", test.input, err)
		}
		if result != test.expected {
			t.Errorf("GetUTFLength(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}