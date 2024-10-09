package main

import("fmt")

func main() {

	nums1 := []int{1, 2, 3}
	fmt.Printf("Слайс #1 %d\n", nums1)
	
	nums2 := []int{4, 5, 6, 7}
	fmt.Printf("Слайс #2 %d\n", nums2)

	fmt.Printf("Join %d\n", Join(nums1, nums2))

}

// Даны 2 слайса целых чисел nums1 и nums2. Создайте функцию Join(nums1, nums2 []int) []int, 
// которя создаст новый слайс емкостью, вмещающей в себя ровно два слайса (ёмкость должна быть равна его длине). 
// Скопируйте в него сначала значения nums1 затем nums2 и верните его.
func Join(nums1, nums2 []int) []int {
    // Создаем новый слайс с длиной и емкостью равными сумме длин nums1 и nums2
	newNums := make([]int, 0, len(nums1)+len(nums2))
 
                                	// вариативный оператор ... - Троеточие после nums2 используется для распаковки элементов слайса nums2, чтобы они были добавлены как отдельные элементы, а не как один слайс
	newNums = append(newNums, nums1...) // добавляем все элементы из nums1
	newNums = append(newNums, nums2...) // добавляем все элементы из nums2
	return newNums
}