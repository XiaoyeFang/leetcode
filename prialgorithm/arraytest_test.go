package prialgorithm

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{3, 2, 3, 1, 5, 0}
	sum := MaxProfit(prices)
	t.Log("sum :=", sum)
}

func TestIsValidSudoku(t *testing.T) {
	nums1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '0'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	//nums = [][]byte{
	//	{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	//}
	valid := IsValidSudoku(nums1)
	t.Log(valid)
}

func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{1, 1, 2, 2, 2, 3, 3}
	RemoveDuplicatesMid(nums)
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{1, 0, 2, 0, 0, 3, 3}
	MoveZeroes(nums)
}

func TestMergeList(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0, 0}
	nums2 := []int{4, 5, 6, 7}

	MergeList(nums1, 3, nums2, len(nums2))
}

func TestRemoveElement(t *testing.T) {
	//nums := []int{2, 1, 2, 3, 2}
	//RemoveElement(nums, 2)
	a := "19"
	b := "19"
	MultiplyLargeNumbers(a, b)
	//lenA := len(a) - 1
	//lenB := len(b) - 1
	//num := int(a[lenA]-'0') * int(b[lenB]-'0')
	//f := num / 10
	//e := num % 10

}
