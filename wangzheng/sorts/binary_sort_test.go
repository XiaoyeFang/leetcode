package sorts

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 6, 8, 9}
	fmt.Println(BinarySearch(nums, 1))
	//
	//start := BinarySearch(nums, 5)
	//end := BinarySearch(nums, 8)
	//
	//fmt.Println(nums[start:end+1])
}

func TestBinarySearchFirst(t *testing.T) {

	nums := []int{1, 2, 3, 4, 6, 8, 9, 9, 9, 9, 9, 9}
	fmt.Println(BinarySearchFirst(nums, 9))
}

func TestBinarySearchLast(t *testing.T) {
	nums := []int{1, 2, 3, 4, 6, 8, 9, 9, 9, 9, 9, 9, 10}
	fmt.Println(BinarySearchLast(nums, 9))
}

func TestBinarySearchBig(t *testing.T) {

	nums := []int{1, 2, 3, 4, 6, 8, 9, 9, 9, 9, 9, 9, 10,11,11}
	fmt.Println(BinarySearchBig(nums, 9))
}