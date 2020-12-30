package main

import (
	"fmt"
	"strings"
)

/**
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."

限制：
0 <= s 的长度 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof
*/

func replaceSpace1(s string) string {
	s = strings.Replace(s, " ", "%20", -1)
	return s
}

func replaceSpace2(s string) string {
	var r []rune
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			r = append(r, '%')
			r = append(r, '2')
			r = append(r, '0')
			continue
		}
		r = append(r, rune(s[i]))
	}

	return string(r)
}

func replaceSpace3(s string) string {
	var str string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			str += "%20"
			continue
		}
		str += string(s[i])
	}
	return str
}

func main() {
	s := replaceSpace3("We are happy.")
	fmt.Println(s)
}
