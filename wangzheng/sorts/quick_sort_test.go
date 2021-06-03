package sorts

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{1, 2, 3, 4, 4, 5, 3, 23, 2, 3, 322, 52}
	QuickSort(nums, 0, len(nums)-1)
	fmt.Printf("nums: %+v \n", nums)
}

func TestGetMid(t *testing.T) {
	fmt.Println(GetMid(1, 2, 3))
	fmt.Println(GetMid(1, 3, 2))
	fmt.Println(GetMid(2, 1, 3))
	fmt.Println(GetMid(2, 3, 1))
	fmt.Println(GetMid(3, 1, 2))
	fmt.Println(GetMid(3, 2, 1))

}

func TestGetNMax(t *testing.T) {
	nums := []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9}
	mid := GetNMax(nums, 0, len(nums)-1, 6)
	fmt.Printf("mid: %+v \n", mid)
}
