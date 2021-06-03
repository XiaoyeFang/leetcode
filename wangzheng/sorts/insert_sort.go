package sorts

//[]int{6,2,5,7,6,3,3,5,2,9,4,1}
//[]int{1,6,2,5,7,6,3,3,5,2,9,4}

func InsertSort(nums []int) {
	n := len(nums) - 1
	for i := 1; i < n; i++ {
		mid := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > mid {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[i] = mid
	}

}
