package main

import(
	"fmt"
)

func main() {

	// Объявляем map
	myMap := make(map[int]int)

    // Добавляем элементы
    myMap[1] = 1000
    myMap[100] = 10000

	fmt.Printf("%d\n", FindMaxKey(myMap))

}

// Напишите функцию FindMaxKey(m map[int]int) int, которая принимает мапу и возвращает значение наибольшего ключа.
// Примечания Например, если передать функции FindMaxKey(m map[int]int) int мапу [10:37 3:19], то она должна вернуть число 10.
func FindMaxKey(m map[int]int) int {
	first := true
	var maxKey int
	for key := range m {
		if first {
			maxKey = key
			first = false
		}
		if key > maxKey {
			maxKey = key
		}
	}
	return maxKey
}