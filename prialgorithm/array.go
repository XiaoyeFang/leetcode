package prialgorithm

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
func MaxProfit(prices []int) int {
	sum := 0
	//只要赚就卖掉
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			sum += prices[i] - prices[i-1]

		}

	}

	return sum
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
