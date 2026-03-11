package main

import (
	"sort"
)

// Medium
// по логике  что бы разбить анаграммы на группы - ключ группы должен содержать в себе конечные символы из слов
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		key := sortString(str)
		//если такой ключ найден - записываем, если нет создастся новый
		m[key] = append(m[key], str)

	}
	result := [][]string{}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	groupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"})
}
