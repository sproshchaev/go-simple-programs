package main

import(
	"fmt"
)

func main() {

	input := map[string]int {
		"qw": 0,
		"qwerty": 1,
        "qwertyu": 2,
	}
	fmt.Printf("%v\n", DeleteLongKeys(input))
}

// Создайте функцию DeleteLongKeys(m map[string]int) map[string]int, которая принимает мапу и возвращает новую мапу, 
// где присутствуют только те ключи, у которых длина не меньше 6 символов.
func DeleteLongKeys(m map[string]int) map[string]int {
	newMap := make(map[string]int)
	for key, value := range m {
		if len(key) > 5 {
			newMap[key] = value
		}
	}
	return newMap
}
