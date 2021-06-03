package sorts

func BinarySearch(nums []int, n int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		if nums[mid] < n {
			low = mid + 1
		} else if nums[mid] > n {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 查找第一个等于给定值的下标
func BinarySearchFirst(nums []int, n int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		if nums[mid] < n {
			low = mid + 1
		} else if nums[mid] > n {
			high = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != n {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// 查找最后一个等于给定值的下标
func BinarySearchLast(nums []int, n int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		if nums[mid] < n {
			low = mid + 1
		} else if nums[mid] > n {
			high = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != n {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

// 查找第一个大于给定值的下标   第一个小于它的同理
func BinarySearchBig(nums []int, n int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		if nums[mid] > n {
			if mid == 0 || nums[mid-1] <= n {
				return mid
			}else {
				high = mid - 1
			}
		} else  {
			low = mid + 1
		}
	}
	return -1
}
