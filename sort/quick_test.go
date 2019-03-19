package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{1, 7, 5, 8, 6, 4, 3, 2, 0}
	QuickSort(nums)
	fmt.Println(nums)

}
