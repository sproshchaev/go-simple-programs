package main

import("fmt")

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := Mix(nums)
	fmt.Println(result) // Вывод: [1 5 2 6 3 7 4 8]

}

// Дан слайс nums, состоящий из 2n элементов в формате [x0,x1,...,xn,y0,y1,...,yn]. 
// Создайте функцию Mix(nums []int) []int, которая вернёт слайс, содержащий значения в следующем порядке: [x0,y0,x1,y1,...,xn,yn].
func Mix(nums []int) []int {
	n := len(nums) / 2 // Так как nums состоит из 2n элементов
	mixed := make([]int, n*2) // Новый слайс длиной 2n

	for i := 0; i < n; i++ {
		mixed[2*i] = nums[i]      // x0, x1, ..., xn
		mixed[2*i+1] = nums[n+i]  // y0, y1, ..., yn
	}

	return mixed
}