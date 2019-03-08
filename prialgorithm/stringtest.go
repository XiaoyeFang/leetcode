package prialgorithm

import (
	"math"
	"strings"
)

//编写一个函数，其作用是将输入的字符串反转过来。
func reverseString(s string) string {
	runes := []rune(s)
	l := len(runes)

	for i, j := 0, l; i < j; i, j = i+1, j-1 {

		runes[i], runes[l-1-i] = runes[l-1-i], runes[i]

	}
	return string(runes)
}

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
func reverse(x int) int {

	sum := 0
	for {
		//不断去最后一位 进行累加累乘
		if x != 0 {
			sum = sum*10 + x%10
			x = x / 10
		} else {
			break
		}
	}

	if x < 0 {
		sum = sum * (-1)
	}
	if sum > math.MaxInt32 || sum < math.MinInt32 {
		return 0
	}

	return sum

}

//字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	length := len(s)
	k := 0
	if length <= 0 {
		return -1
	}

	repeated := false
	for i := 0; i < length; i++ {
		repeated = false
		for j := 0; j < length; j++ {
			if j != i && s[j] == s[i] {
				repeated = true
				break
			}
			k = i
		}
		if !repeated {
			return k
		}

	}

	return -1

}

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mp1 := make(map[rune]int)
	mp2 := make(map[rune]int)
	for _, v := range s {
		if mp1[v] != 0 {
			mp1[v]++
		} else {
			mp1[v] = 1
		}
	}
	for _, v := range t {
		if mp2[v] != 0 {
			mp2[v]++
		} else {
			mp2[v] = 1
		}
	}
	for k, v := range mp1 {
		if mp2[k] != v {
			return false
		}
	}

	return true

}

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

//你们这题太欺负人了你知道吗，只考虑字母和数字字符你们还把标点符号算进去。。。 555555555555555
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	runes := []rune(s)

	low := 0
	high := len(s) - 1

	for {

		if low < high {
			for {
				if low < high && !isAlphanum(runes[low]) {
					//从前往后遍历，既不是字母又不数字，则略过
					low++
				} else {
					break
				}
			}

			for {
				if low < high && !isAlphanum(runes[high]) {
					//从后往前遍历，既不是字母又不数字，则略过
					high--
				} else {
					break
				}
			}

			//如果对称位置的不相等
			if runes[low] != runes[high] {
				return false
			}
			low++
			high--
		} else {
			break
		}
	}

	return true

}

func isAlphanum(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')

}

//请你来实现一个 atoi 函数，使其能将字符串转换成整数。
func myAtoi(str string) int {
	//     这个真尼玛难啊，我连go里的atoi方法里的都用上了还没解开，将来我如果看到这个，一定要穿越回来告诉我怎么做
	if str == "" {
		return 0
	}
	runes := []rune(str)
	var sum = 0
	var negative = false
	if !isNum(runes[0]) && runes[0] != '-' && runes[0] != ' ' {
		return 0
	}
	for _, v := range runes {
		if v == '-' {
			negative = true
		}
		if isNum(v) {
			sum = sum*10 + int(v-48)
		}
	}

	if negative {
		sum = sum * -1
	}

	return sum
}

func isNum(c rune) bool {
	return '0' <= c && c <= '9'

}

//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {

		for j := 0; j < len(needle); j++ {

			if haystack[i] == needle[j] {
				if i+len(needle) > len(haystack) {
					return -1
				}

				if haystack[i:i+len(needle)] == needle {
					return i
				}

			}
		}

	}

	return -1
}

//编写一个函数来查找字符串数组中的最长公共前缀。 如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		// 找出S1与Si间的最长公共字符串
		// indexOf：当存在串时返回>0；不存串时返回-1
		// 只要不存在串，就缩减串的规模，再进行查找
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}

		}
	}
	return prefix
}
