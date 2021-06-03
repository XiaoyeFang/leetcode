package sort

import (
	"fmt"
	"sort"
)

type Nums struct {
	nums []int
}

func NewNums() *Nums {

	return &Nums{make([]int, 0)}
}

func (n *Nums) Selelction() {
	for i := 0; i < len(n.nums); i++ {

		min := i
		for j := i + 1; j < len(n.nums); j++ {
			if n.nums[j] < n.nums[min] {
				min = j
			}

		}
		n.nums[i], n.nums[min] = n.nums[min], n.nums[i]

	}

	fmt.Println(n.nums)

}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	// 排序后 依次用左边界与其他右边界比较，当超过左边界时，下角标++  箭数量++
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	var min = 1
	//给一个最初起点  一般是第一个气球的最右边
	var right = points[0][1]
	for _, p := range points {
		if p[0] > right {
			right = p[1]
			min++
		}
	}
	return min
}

func minMoves(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var min int
	for _, n := range nums {
		min = n - nums[0]
	}
	return min
}
