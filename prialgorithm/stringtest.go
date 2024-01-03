package prialgorithm

import (
	"math"
	"strings"
)

// 编写一个函数，其作用是将输入的字符串反转过来。
func reverseString(s string) string {
	runes := []rune(s)
	l := len(runes)

	for i, j := 0, l; i < j; i, j = i+1, j-1 {

		runes[i], runes[l-1-i] = runes[l-1-i], runes[i]

	}
	return string(runes)
}

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
func Reverse(x int) int {

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

// 字符串中的第一个唯一字符
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

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。
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

// 你们这题太欺负人了你知道吗，只考虑字母和数字字符你们还把标点符号算进去。。。 555555555555555
func isPalindString(s string) bool {
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
//func MyAtoi(str string) int {
//这个真尼玛难啊，我连go里的atoi方法里的都用上了还没解开，将来我如果看到这个，一定要穿越回来告诉我怎么做
//兄弟 我回来了 我来告诉你
/*
请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
 核心：result=result∗10+int(s[i]−′0′)
*/
func MyAtoi(s string) int {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}
	// 定义符号位
	sign, i := 1, 0

	if s[i] == '+' || s[i] == '-' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	num := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}

		digit := int(s[i] - '0')

		if num > math.MaxInt32/10 || (num == math.MaxInt32/10 && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		num = num*10 + digit
	}

	return num * sign
}

func isNum(c rune) bool {
	return '0' <= c && c <= '9'

}

// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1
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

// 编写一个函数来查找字符串数组中的最长公共前缀。 如果不存在公共前缀，返回空字符串 ""。
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

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
划定左右边界，判定有重复
*/
func LengthOfLongestSubstring(s string) int {
	left := 0
	res := 0
	tmp := 0
	for i := 0; i < len(s); i++ {
		tmp = i - left + 1
		for j := left; j < i; j++ {
			if s[i] == s[j] {
				tmp = i - left
				left = j + 1
				break
			}
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}

/*
不重复 map[元素]出现次数
子串 滑动窗口
需要理解两个点，
1.意味着上一次退出是有重复的元素
*/
func LengthOfLongestSubstringTwo(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1]) //1
		}
		for rk+1 < n && m[s[rk+1]] == 0 { //2
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

//1.初始化三个值
//map 存储出现过的值
//res 存储出现过的最长子串 start 右指针
// 两层for 从头遍历到尾  遍历最长的子串，每次新加入子串的元素都要在map中查找是否出现过，出现过则换个起始点遍历寻找子串

func LengthOfLongestSubstringThree(s string) int {
	var start = -1
	var res int
	var m = map[uint8]int{}
	for i := 0; i < len(s); i++ {
		//
		if i != 0 {
			delete(m, s[i])
		}

		for start+1 < len(s) && m[s[start+1]] == 0 {
			m[s[start+1]]++
			start++
		}

		res = max(res, start-i+1)
	}
	return res
}

/*
给你一个字符串 s，找到 s 中最长的回文子串。
解题思路 首先要明白怎么去找回文子串，遍历时，每到一个下标，先去左右扩散式寻找回文串 此时有可能出现两种情况 有中心的回文串和无中心的回文串
*/
func LongestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		len1 := ExpandAroundCenter(s, i, i)   // 以一个字符为中心的回文串
		len2 := ExpandAroundCenter(s, i, i+1) // 以两个字符之间为中心的回文串
		maxLen := max(len1, len2)
		if maxLen > end-start { // 如果当前获取的回文子串比之前的更长 去更新start end
			start = i - (maxLen-1)/2 //根据返回的子串长度去更新
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

func ExpandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] { //向左右扩散判断是否是回文字符串
		left--
		right++
	}
	return right - left - 1 //当前起始点已经找到的最大回文字符串 为什么是-1 因为循环最终的left和right不是回文对应的
}
