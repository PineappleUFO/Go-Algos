package main

import (
	"fmt"
)

/*
 * Берем самую короткую строку - т.к. общий префикс не может быть длинее самой короткой строки
 * Идем по буквам короткого слова
 * Для каждой буквы проверяем по словам из str
 * Если буква не совпала возвращаем накопленый префикс
 * Сложность O(M x N) M - длина самого короткого слова, N- колво слов
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	shortest := strs[0]
	prefix := ""
	for _, word := range strs {
		if len(shortest) > len(word) {
			shortest = word
		}

	}
	for ind := 0; ind < len(shortest); ind++ {
		for _, word := range strs {
			if word[ind] != shortest[ind] {
				return prefix
			}
		}
		prefix += string(shortest[ind])
	}

	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"dance", "dag", "danger", "damage"}))
}
