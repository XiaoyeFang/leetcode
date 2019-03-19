package sort

import "fmt"

/*
冒泡排序 把nums当做竖直的水柱，气泡逐渐上升,当前值不停地与最大值交换位置
 */
func (n *Nums) Bubble() {
	len := len(n.nums)
	for i := 0; i < len-1; i++ {
		for j := i + 1; i < len-1; j++ {

			if n.nums[i] < n.nums[j] {
				n.nums[i], n.nums[j] = n.nums[j], n.nums[i]
			}
		}

	}
	fmt.Println(n.nums)
}
