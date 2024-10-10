package main

import(
	"fmt"
)

func main() {

	// Создаем мапу
	myMap := map[string]string{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	}

	fmt.Printf("%v\n", SwapKeysAndValues(myMap))

}

// Напишите функцию SwapKeysAndValues(m map[string]string) map[string]string, которая принимает мапу и возвращает новую мапу, где ключи и значения поменялись местами.
func SwapKeysAndValues(m map[string]string) map[string]string {
	return nil
}