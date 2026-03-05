package main

import "fmt"

// фишка в перевернутом словаре, в seen в индексе храним значение из nums | в значении храним индекс из nums
// итерируемся по входящему массиву и проверям по ключу словаря есть ли (target-key) в словаре
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for ind, key := range nums {
		if _, ok := seen[target-key]; ok {
			return []int{seen[target-key], ind}
		}
		seen[key] = ind
	}
	return []int{0, 0}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 6
	fmt.Println(twoSum(nums, target))
}
