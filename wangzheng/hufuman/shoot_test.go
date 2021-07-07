package hufuman

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

type person_api interface {
	say()
}
type person_2api interface {
	person_api
	run()
}

type lawyer struct {
	Name string
}

func (this lawyer) say() {
	fmt.Printf("i am lawyer %s\n", this.Name)
}

type teacher struct {
	Name string
}

func (this teacher) say() {
	fmt.Printf("i am teacher %s\n", this.Name)
}
func (this teacher) run() {
	fmt.Printf("teacher %s is running\n", this.Name)
}

func Say(p person_api) {
	p.say()
}

func TestMinShoot(t *testing.T) {
	//aims :=[][]int{{1,6},{2,3},{4,5},{9,10}}
	//fmt.Println(MinShoot(aims))
	//nums :=[]int32{1}
	//for i := 0; i < 10; i++ {
	//	fmt.Println(generateRandomNumber(0, 20, 10))
	//}
	var a1 person_api = lawyer{"sccot"}
	//a1.say()
	//var a2 person_2api = teacher{"wiliam"}
	//a2.say()
	//a2.run()
	Say(a1)

}

func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}
