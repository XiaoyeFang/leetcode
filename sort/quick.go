package sort

import "fmt"

func QuickSort(nums []int) {
	if len(nums) < 1 {
		return
	}
	mid := nums[0] //6
	i := 1
	head, tail := 0, len(nums)-1
	for head < tail {
		fmt.Println(nums)
		if nums[i] < mid {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++

		} else {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--

		}
	}
	
	nums[head] = mid
	QuickSort(nums[:head])
	QuickSort(nums[head+1:])
}
