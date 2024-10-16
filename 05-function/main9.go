package main

import (
	"fmt"
	"sort"
)

func main() {
	input := "cab"
	permutations := Permutations(input)

	fmt.Printf("Перестановки строки \"%s\": %v\n", input, permutations)
}

// Напишите функцию Permutations(input string) []string, которая принимает строку и выводит все перестановки её символов в алфавитном порядке.
// Примечания Например, если передать функции Permutations(input string) []string строку "abc" или "cab", то она должна вернуть [abc acb bac bca cab cba]. 
// Для сортировки объектов используйте пакет sort.
func Permutations(input string) []string {
	var result []string
	permutate([]rune(input), 0, &result)
	sort.Strings(result) // Сортируем полученные перестановки
	return result
}

// permutate рекурсивно генерирует перестановки
func permutate(runes []rune, index int, result *[]string) {
	if index == len(runes)-1 {
		*result = append(*result, string(runes))
		return
	}

	// Используем множество для предотвращения дублирования перестановок
	seen := make(map[rune]bool)

	for i := index; i < len(runes); i++ {
		if seen[runes[i]] {
			continue // Пропускаем уже обработанные символы
		}
		seen[runes[i]] = true

		// Меняем местами текущий символ с символом на позиции index
		runes[index], runes[i] = runes[i], runes[index]
		permutate(runes, index+1, result) // Рекурсивный вызов
		runes[index], runes[i] = runes[i], runes[index] // Возврат к предыдущему состоянию
	}
}
