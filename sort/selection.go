package sort

import "fmt"

type Nums struct {
	nums []int
}

func NewNums() *Nums {

	return &Nums{make([]int, 0)}
}

func (n *Nums) Selelction() {
	for i := 0; i < len(n.nums); i++ {

		min := i
		for j := i + 1; j < len(n.nums); j++ {
			if n.nums[j] < n.nums[min] {
				min = j
			}

		}
		n.nums[i], n.nums[min] = n.nums[min], n.nums[i]

	}

	fmt.Println(n.nums)


}
