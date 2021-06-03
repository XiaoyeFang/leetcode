package sorts

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	nums := []int{1, 2, 3, 4, 4, 5, 3, 23, 2, 3, 322, 52}
	InsertSort(nums)
	fmt.Println("nums: ", nums)
}
