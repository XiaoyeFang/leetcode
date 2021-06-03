package sorts

func QuickSort(nums []int, low, high int) {
	if low < high {
		i := partition(nums, low, high)
		QuickSort(nums, low, i-1)
		QuickSort(nums, i+1, high)
	}

}

/* 优化方案
1 三数取中位数
2 随机取基准数
*/
//取中位数方法如下
func GetMid(a, b, c int) int {
	if (a-b)*(b-c) > 0 {
		return b
	}
	if (a-c)*(c-b) > 0 {
		return c
	}
	return a
}

func partition(nums []int, low, high int) int {
	i := low + 1
	j := high

	for i < j {
		if nums[i] > nums[low] {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	if nums[i] > nums[low] {
		i--
	}
	nums[i], nums[low] = nums[low], nums[i]
	return i
}

//快速排序如何取第N大的数据

func GetMaxSum(a, b, c, n int) int {
	if (a-b)*(b-c) > 0 {
		return b
	}
	if (a-c)*(c-b) > 0 {
		return c
	}
	return a
}

func GetNMax(nums []int, low, high, k int) int {
	i := partitionMax(nums, low, high)

	for i != k-1 {
		// 此时划分好三个区，[:i] [i] [i:]
		if i < k-1 {

			i = partitionMax(nums, i+1, len(nums)-1)

		} else {
			i = partitionMax(nums, 0, i-1)

		}
	}
	return nums[i]
}
func partitionMax(nums []int, low, high int) int {
	i := low + 1
	j := high

	for i < j {
		if nums[i] > nums[low] {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	if nums[i] > nums[low] {
		i--
	}
	nums[i], nums[low] = nums[low], nums[i]
	return i
}
