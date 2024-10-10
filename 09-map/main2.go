package main

import(
	"fmt"
)

func main() {

    // Создание мап-ы
	myMap := map[int]int {
		1: 10,
		2: 20,
		3: 30,
	}

	fmt.Printf("%d\n", SumOfValuesInMap(myMap))

}

// Напишите функцию SumOfValuesInMap(m map[int]int) int, которая возвращает сумму значений в мапе.
// Примечания: Например, если передать функции SumOfValuesInMap(m map[int]int) int мапу [10:38 3:19], то она должна вернуть число 57.
func SumOfValuesInMap(m map[int]int) int {
	var sumValue int
    // Перебор только значений 
	for _, value := range m {
		sumValue = sumValue + value
	}
	return sumValue
}
