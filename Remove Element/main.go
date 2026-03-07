package main

import "fmt"
//два указателя — k указывает на позицию для следующего «нужного» элемента. Проходим по массиву: если элемент ≠ val, записываем его на позицию k и сдвигаем k вперёд.
//Сложность: O(n) по времени, O(1) по памяти.
func removeElement(nums []int, val int) int {

	k := 0
	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}
	return k
}

func main() {
	fmt.Println(removeElement([]int{1, 1, 2, 3, 4}, 1))
}
