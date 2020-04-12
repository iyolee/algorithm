package string

// // 编辑距离又称Levenshtein距离，是指两个字符串之间，由一个转成另一个所需的最少编辑操作次数。

// // replaceWight 用来表示替换操作与插入删除操作相比的倍数
// func edit(s1, s2 string, replaceWight int) int {
// 	if s1 == "" || s2 == "" {
// 		return
// 	}

// 	if s1 == "" {
// 		return len(s2)
// 	}

// 	if s2 == "" {
// 		return len(s1)
// 	}

// 	len1, len2 := len(s1), len(s2)

// 	// 申请二维数组来存储中间的计算结果
// 	D := make([][]int, len1 + 1)

// 	for i, _ := range D {
// 		D[i] = make([]int, len2 + 1)
// 	}

// 	for i:= 0; i < len1 + 1; i++ {
// 		D[i][0] = i
// 	}
// }
