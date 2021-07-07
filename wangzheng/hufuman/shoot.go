package hufuman

import (
	"fmt"
	"sort"
)

func MinShoot(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	// 排序后 依次用左边界与其他右边界比较，当超过左边界时，下角标++  箭数量—++
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	fmt.Println(points)
	var min =1
	//给一个最初起点  一般是第一个气球的最右边
	var right = points[0][1]
	for _, p := range points {
		if p[0] > right {
			right = p[1]
			min++
		}
	}
	return min

}

//func ()  {
//
//}