package prialgorithm

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	sum := Reverse(-4432)
	t.Log(sum)
}

func TestMyAtoi(t *testing.T) {
	//str := " -42"
	//num := MyAtoi(str)
	//t.Logf("%b \n", fufeibonachi(3))
	testCase := []int{
		1, 2, 4,
	}
	res := rob(testCase)
	fmt.Println(res)
	fmt.Println(int(0x00000011))
}
func fufeibonachi(i int) int{
	if i==1 ||i ==2{
		return 1
	}

	return fufeibonachi(i-1)+fufeibonachi(i-2)
}

func rob(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return maxInt(nums[0], nums[1])
	}
	return maxInt(rob(nums[:len(nums)-3])+nums[len(nums)-1], rob(nums[:len(nums)-2]))
}

func maxInt(args ...int) int {
	maxInt := args[0]
	for index, a := range args {
		if a > maxInt {
			maxInt = args[index]
		}
	}
	return maxInt
}