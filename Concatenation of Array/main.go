package main

import "fmt"

// конкатенация массивов, изи
func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

func main() {
	fmt.Println(getConcatenation([]int{1, 4, 1, 2}))
}
