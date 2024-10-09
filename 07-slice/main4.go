package main

import("fmt")

func main() {

	nums1 := make([]int, 2)
	nums1[0] = 1
	nums1[1] = 2
	fmt.Printf("Слайс #1 %d\n", nums1)
	
	nums2 := make([]int, 2)
	nums2[0] = 1
	nums2[1] = 2
	fmt.Printf("Слайс #2 %d\n", nums2)

}

// Даны 2 слайса целых чисел nums1 и nums2. Создайте функцию Join(nums1, nums2 []int) []int, 
// которя создаст новый слайс емкостью, вмещающей в себя ровно два слайса (ёмкость должна быть равна его длине). 
// Скопируйте в него сначала значения nums1 затем nums2 и верните его.
