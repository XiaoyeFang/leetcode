package sort

import "testing"

func TestNums_Bubble(t *testing.T) {
	n := NewNums()
	n.nums = append(n.nums, 1,7,5,8,6,4,3,2,0)
	n.Selelction()
}