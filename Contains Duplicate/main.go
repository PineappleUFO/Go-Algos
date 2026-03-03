package main

import "fmt"

func hasDuplicate(nums []int) bool {
	set := make(map[int]struct{}, len(nums)) //создаем пустой map с ограничением для лучшей алокации памяти
	for _, num := range nums {               //итерируемся по входящему массиву
		if _, ok := set[num]; ok { // set:1 свойство значение - нам не нужно, 2 свойство true|false в зависимости от наличия ключа
			return true
		}
		set[num] = struct{}{} ////записываем ключ без значения, нам оно не нужно
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 5, 5}
	if hasDuplicate(nums) {
		fmt.Println("Найден дубликат")
	}
}
