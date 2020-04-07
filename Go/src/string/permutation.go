package string

// 字符串全排列

import (
	"fmt"
)

// Permutation Permutation
func Permutation(str []rune, start int) {
	if str == nil {
		return
	}

	if (start == len(str) - 1) {
		fmt.Println(string(str))
	} else {
		for i := start; i < len(str); i++ {
			if (!isDuplicate(str, start, i)) {
				continue
			}
			// 交换start与i所在位置的字符
			SwapRune(str, start, i)
			// 固定第一个字符，对剩余的字符进行全排列
			Permutation(str, start + 1)
			// 还原start与i所在位置的字符
			SwapRune(str, start, i)
		}
	}
}

// SwapRune SwapRune
func SwapRune(data []rune, x, y int) {
	tmp := data[x]
	data[x] = data[y]
	data[y] = tmp
}

func isDuplicate(str []rune, begin, end int) bool {
	for i := begin; i < end; i++ {
		if (str[i] == str[end]) {
			return false
		}
	}
	return true
}

// PermutationStr PermutationStr
func PermutationStr(s string) {
	Permutation([]rune(s), 0)
}