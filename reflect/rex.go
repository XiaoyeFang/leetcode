package reflect

import (
	"fmt"
	"reflect"
	"sync/atomic"
)

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch v.(type) {

	case *student, student:
		fmt.Printf("%s \n", reflect.ValueOf(v))
		//fmt.Println("111111")
	}
}

func (s *student) Reflectex() string {
	//s := &student{
	//	"sss",
	//}
	//zhoujielun(s)
	return fmt.Sprintf("%v", s)
}

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %s", *p)
}

var value int32

func SetValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, v + delta) {
			fmt.Println("aaaaa ")
			break
		}
	}
}
