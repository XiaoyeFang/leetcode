package sort

import "testing"

func TestNums_Selelction(t *testing.T) {
	n := NewNums()
	n.nums = append(n.nums, 0,2, 1, 3)
	n.Selelction()

}
