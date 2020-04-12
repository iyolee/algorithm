package string

import (
	"bytes"
	"errors"
)

// 根据两个文件的绝对路径算出其相对路径
// a="/qihoo/app/a/b/c/d/new.c"
// b="/qihoo/app/1/2/test.c"
// 那么b相对于a的相对路径是"../../../../1/2/test.c"

func GetRelativePath(path1, path2 string) string {
	if path1 == "" || path2 == "" {
		defer errors.New("参数不合法")
	}

	arr1, arr2 := []rune(path1), []rune(path2)

	// 用来指向两个路径中不同目录的起始路径
	diff1, diff2 := 0, 0

	len1, len2 := len(arr1), len(arr2)
	relativePath := new(bytes.Buffer)

	for i, j := 0, 0; i < len1 && j < len2; {
		if arr1[i] == arr2[j] {
			if arr1[i] == '/' {
				diff1, diff2 = i, j
			}
			i++
			j++
		} else {
			// 跳过目录分隔符 /
			diff1++
			for diff1 < len1 {
				if arr1[diff1] == '/' {
					relativePath.WriteString("../")
				}
				diff1++
			}
			diff2++
			relativePath.WriteString(string(arr2[diff2:]))
			break
		}
	}
	return relativePath.String()
}
