package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(AreAnagrams("listen", "silent"))        // true
	fmt.Println(AreAnagrams("triangle", "integral"))    // true
	fmt.Println(AreAnagrams("apple", "pale"))           // false
	fmt.Println(AreAnagrams("Dormitory", "Dirty room")) // true
}

// Напишите функцию AreAnagrams(str1, str2 string) bool, которая проверяет, являются ли две заданные строки анаграммами 
// (то есть состоят ли они из одних и тех же символов, хотя и в разном порядке). Не учитывайте регистр символов.
// Примечания: Используйте встроенные функции сортировки и сравнения строк.
func AreAnagrams(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	sort.Slice(runeStr1, func(i, j int) bool {
		return runeStr1[i] < runeStr1[j]
	})
	sort.Slice(runeStr2, func(i, j int) bool {
		return runeStr2[i] < runeStr2[j]
	})

	return string(runeStr1) == string(runeStr2)
}
