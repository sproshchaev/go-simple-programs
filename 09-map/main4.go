package main

import(
	"fmt"
)

func main() {
	input := []string{"A", "A", "B", "C"}
	// если возвращается map, то используем %v
	fmt.Printf("%v\n", CountingSort(input))
}

// Для этого напишите функцию CountingSort(contacts []string) map[string]int, 
// которая принимает слайс строк и возвращает мапу, где ключ - это элемент слайса, 
// а значение - количество раз, сколько элемент встречается в слайсе.
func CountingSort(contacts []string) map[string]int {
	newMap := make(map[string]int)
	for _, contact := range contacts  {
		value, exists := newMap[contact]
		if exists {
			newMap[contact] = value + 1
		} else {
			newMap[contact] = 1
		}
	}
	return newMap
}
