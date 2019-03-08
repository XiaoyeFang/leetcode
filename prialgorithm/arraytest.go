package prialgorithm

import "fmt"

/*
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
*/
func RemoveDuplicates(nums []int) int {
	i := 0
	var j int
	for {
		if i >= len(nums)-1 {
			break
		}

		for j = i + 1; j < len(nums) && nums[i] == nums[j]; j++ {

		}
		nums = append(nums[:i+1], nums[j:]...)

		i++
	}

	return len(nums)
}

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

贪心算法 跟后面的动态规划股票问题结合理解
*/
//func MaxProfit(prices []int) int {
//	sum := 0
//	//只要赚就卖掉
//	for i := 1; i < len(prices); i++ {
//		if prices[i-1] < prices[i] {
//			sum += prices[i] - prices[i-1]
//
//		}
//
//	}
//
//	return sum
//}

func MaxProfit(prices []int) int {

	//只要赚就卖掉
	for i := 0; i < len(prices)-1; i++ {
		sum := 0
		//for j := i + 1; j < len(prices); j++ {
		if prices[i+1] > prices[i] {
			sum += prices[i+1] - prices[i]
			continue
		}

		//}
		fmt.Println(sum)

	}

	return 0
}

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
*/
func Rotate(nums []int, k int) {
	//简化成了每次向右平移一位，重复k次
	length := len(nums)
	// 右移newk + n * length个位置，和右移newk个位置效果是一样的
	newk := k % length
	temp := 0
	for i := 0; i < newk; i++ {
		temp = nums[length-1]
		for j := length - 2; j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		nums[0] = temp
	}
}

/*
给定一个整数数组，判断是否存在重复元素。
*/
func ContainsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i] == nums[j] {
				return true
			}
		}

	}

	return false
}

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	num := 0

	for i := 0; i < len(nums); i++ {
		num ^= nums[i]

	}

	return num
}

//给定两个数组，编写一个函数来计算它们的交集。
func Intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int) //key 是 nums1中的不重复元素，value是元素在nums1中出现的次数
	res := []int{}
	for _, v := range nums1 {
		m[v] += 1
	}
	for _, v := range nums2 {
		if m[v] > 0 {
			res = append(res, v)
			m[v]--
		}
	}
	return res
}

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
func PlusOne(digits []int) []int {

	index := len(digits) - 1

	add := add1(digits[index]) //add的作用就是判断是进位还是加一

	extend := false //如果全是0，说明要扩展数组，则为true

	for {
		if add == 1 {
			digits[index] = digits[index] + add
			break
		} else {
			digits[index] = 0 //进位，置0
			index--
		}
		if index <= 0 && digits[0] == 0 { //经过计算后为0，说明是进位
			extend = true
			break
		}
		add = add1(digits[index]) //下一位判断是否进位
	}
	if extend {
		ndts := make([]int, len(digits)+1)

		ndts[0] = 1
		for i := 1; i < len(digits)+1; i++ {
			ndts[i] = 0
		}
		return ndts
	} else {

	}
	return digits
}

func add1(digit int) int {
	if digit == 9 {
		return 0
	} else {
		return 1
	}
}

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
func moveZeroes(nums []int) {

	a := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[a], nums[i] = nums[i], nums[a]
			a++
		}
	}
}

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
func TwoSum(nums []int, target int) []int {
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {

		for j := 1; j != i && j < len(nums); j++ {
			if nums[i]+nums[j] == target {

				result = append(result, i, j)
				return result

			}

		}
	}

	return result
}

//判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可

func IsValidSudoku(board [][]byte) bool {

	//创建三个二维数组 依次比较横 纵 方针中是否存在相同数字

	//raw := make([][]byte, 9, 9)
	//col := make([][]byte, 9, 9)
	//cell := make([][]byte, 9, 9)
	//
	//for i := 0; i < len(board); i++ {
	//	for j := 0; j < len(board[i]); j++ {
	//
	//		if board[i][j] == '.' {
	//			continue
	//		}
	//
	//		if isContains(board[i][j], raw[i]) {
	//			fmt.Println("iiiiiii", i)
	//			return false
	//		}
	//		if isContains(board[i][j], col[j]) {
	//			fmt.Println("jjjjjjj", j)
	//			return false
	//		}
	//
	//		//	计算在第几个方阵里
	//		g := i/3*3 + j/3
	//		println("g====", g)
	//		if isContains(board[i][j], cell[g]) {
	//
	//			fmt.Printf("i %v,j %v,board[i][j] %v,raw[i] %v  g %v \n", i, j, board[i][j], cell[g], g)
	//			return false
	//		}
	//
	//		raw[i] = append(raw[i], board[i][j])
	//		col[j] = append(col[j], board[i][j])
	//		cell[g] = append(cell[g], board[i][j])
	//
	//	}
	//	fmt.Printf("cell %v \n", cell)
	//
	//	//fmt.Printf("raw[i] %v \n", raw[i])
	//}
	////fmt.Printf("col[j] %v \n", col)
	//return true

	rowMask, colMask, areaMask := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}

	for row := range board {
		for col := range board[row] {

			if board[row][col] == '.' {
				continue
			}

			digit := board[row][col] - '0' - 1
			area := (row/3)*3 + col/3

			if rowMask[row][digit] || colMask[col][digit] || areaMask[area][digit] {
				return false
			}

			rowMask[row][digit] = true
			colMask[col][digit] = true
			areaMask[area][digit] = true
		}
	}

	return true
}

func isContains(b byte, nums []byte) bool {

	for _, v := range nums {
		if b == v {
			fmt.Printf("b %v ,v,%v \n", string(b), string(v))
			return true
		}
	}
	return false
}

//给定一个 n × n 的二维矩阵表示一个图像。将图像顺时针旋转 90 度。
func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {

			matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]

		}

	}

	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {

			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]

		}

	}
}
