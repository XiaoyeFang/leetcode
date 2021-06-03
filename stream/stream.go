package stream

import (
"encoding/json"
"fmt"
"math"
)

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}



func StreamMap() {

	s1 := Student{1, "张三", 10, 55}
	s2 := Student{2, "李四", 9, 60}
	s3 := Student{3, "王五", 9, 77}
	s4 := Student{4, "赵六", 11, 99}
	s5 := Student{5, "马七", 11, 100}
	s6 := Student{6, "王八", 20, 80}

	arr := make([]interface{}, 0, 0)
	arr = append(arr, s1, s2, s3, s4, s5, s6)

	r1 := make([]Student, len(arr))
	// 寻找所有年龄>10的学生
	Stream(arr).Filter(func(each interface{}) bool {
		return each.(Student).Age > 10
	}).Collect(&r1)


	fmt.Printf("r1 所有年龄>10的学生: %v", r1) //r1 所有年龄>10的学生: [{4 赵六 11 99} {5 马七 11 100} {6 王八 20 80}]
	fmt.Println()

	// 是否有名字叫沙雕的学生?
	r2 := Stream(arr).AnyMatch(func(each interface{}) bool {
		return each.(Student).Name == "沙雕"
	})

	fmt.Printf("r2 是否有名字叫沙雕的学生?: %v", r2) //r2 是否有名字叫沙雕的学生?: false
	fmt.Println()

	// 年龄为11的学生的个数
	r3 := Stream(arr).Filter(func(each interface{}) bool {
		return each.(Student).Age == 11
	}).Count()

	fmt.Printf("r3 年龄为11的学生的个数: %v", r3) //r3 年龄为11的学生的个数: 2
	fmt.Println()

	// 把学生按年龄分组
	r4 := make(map[int64][]Student)
	Stream(arr).GroupByInt(func(each interface{}) int64 {
		return int64(each.(Student).Age)
	}, &r4)

	fmt.Printf("r4 把学生按年龄分组: %v", r4) //r4 把学生按年龄分组: map[9:[{2 李四 9 60} {3 王五 9 77}] 10:[{1 张三 10 55}] 11:[{4 赵六 11 99} {5 马七 11 100}] 20:[{6 王八 20 80}]]
	fmt.Println()

	// 数组去重
	r5 := Stream(1, 2, 3, 4, 5, 1, 2, 3).Distinct()
	fmt.Printf("r5 去重[1, 2, 3, 4, 5, 1, 2, 3]: %v", r5) //r5 去重[1, 2, 3, 4, 5, 1, 2, 3]: [1 2 3 4 5]
	fmt.Println()

	// 找出id为2的学生,打印出年龄
	if r6, ok := Stream(arr).Filter(func(each interface{}) bool {
		return each.(Student).Id == 222
	}).FindAny(); ok {
		fmt.Printf("r6 找出id为2的学生,打印出年龄: %v", r6.(Student).Age)
		fmt.Println()
	} else {
		fmt.Println("no such element")
	}

	r7 := make([]Student, len(arr))
	// 把id>3的学生年龄+10
	Stream(arr).Filter(func(each interface{}) bool {
		return each.(Student).Id > 3
	}).Map(func(each interface{}) interface{} {
		student := each.(Student)
		student.Age = student.Age + 10
		return student
	}).Collect(&r7)
	fmt.Printf("r7 把id>3的学生年龄+10: %v", r7) //r7 把id>3的学生年龄+10: [{4 赵六 21 99} {5 马七 21 100} {6 王八 30 80}]
	fmt.Println()

	// 求所有学生年龄的和
	r8 := Stream(arr).Sum(func(each interface{}) interface{} {
		return each.(Student).Age
	})
	fmt.Printf("r8 求所有学生年龄的和: %v", r8) //r8 求所有学生年龄的和: 70
	fmt.Println()

	// 求所有学生总分和
	r9 := Stream(arr).Reduce(0, func(pre interface{}, cur interface{}) interface{} {
		return pre.(int) + cur.(Student).Score
	})
	fmt.Printf("r9 求所有学生总分和: %v", r9) //r9 求所有学生总分和: 471
	fmt.Println()

	// 将id为偶数的学生的分数+10, 再求和
	r10 := Stream(arr).Filter(func(each interface{}) bool {
		return each.(Student).Id&1 == 0
	}).Map(func(each interface{}) interface{} {
		student := each.(Student)
		student.Age = student.Score + 10
		return student
	}).Reduce(0, func(pre interface{}, cur interface{}) interface{} {
		return pre.(int) + cur.(Student).Score
	})

	fmt.Printf("r10 将id为偶数的学生的分数+10, 再求和: %v", r10) //r10 将id为偶数的学生的分数+10, 再求和: 239
	fmt.Println()

	// 求所有学生的平均分数
	r11 := Stream(arr).Average(func(each interface{}) interface{} {
		return each.(Student).Score
	})

	fmt.Printf("r11 求所有学生的平均分数: %v", r11) //r11 求所有学生的平均分数: 78.5
	fmt.Println()

	// 求所有学生的最高分
	r12 := Stream(arr).Max(func(each interface{}) interface{} {
		return each.(Student).Score
	})

	fmt.Printf("r12 求所有学生的最高分: %v", r12) //r12 求所有学生的最高分: 100
	fmt.Println()

	// 求所有学生的最低分
	r13 := Stream(arr).Min(func(each interface{}) interface{} {
		return each.(Student).Score
	})

	fmt.Printf("r13 求所有学生的最低分: %v", r13)  //r13 求所有学生的最低分: 55
	fmt.Println()
}

func If(condition bool, trueResult interface{}, falseResult interface{}) interface{} {
	if condition {
		return trueResult
	}
	return falseResult
}

type stream struct {
	list []interface{}
}

func Stream(arrs ...interface{}) *stream {

	st := new(stream)

	if len(arrs) > 0 {
		if x, ok := arrs[0].([]interface{}); ok {
			st.list = make([]interface{}, len(x))
			copy(st.list, x)
		} else {
			st.list = make([]interface{}, len(arrs))
			copy(st.list, arrs)
		}
	}

	return st
}

func (s *stream) Filter(fn func(each interface{}) bool) *stream {
	list := make([]interface{}, 0, len(s.list))
	for _, x := range s.list {
		if fn(x) {
			list = append(list, x)
		}
	}
	s.list = list
	return s
}

func (s *stream) ForEach(fn func(each interface{})) {
	list := s.list
	for _, x := range list {
		fn(x)
	}
}

func (s *stream) Collect(r interface{}) {
	bytes, _ := json.Marshal(s.list)
	json.Unmarshal(bytes, &r)
}

func (s *stream) FindAny() (interface{}, bool) {
	if len(s.list) > 0 {
		return s.list[0], true
	}
	return nil, false
}

func (s *stream) AnyMatch(fn func(each interface{}) bool) bool {
	for _, x := range s.list {
		if fn(x) {
			return true
		}
	}
	return false
}

func (s *stream) Map(fn func(each interface{}) interface{}) *stream {
	for i, x := range s.list {
		s.list[i] = fn(x)
	}
	return s
}

func (s *stream) Count() int {
	return len(s.list)
}

func (s *stream) Distinct() []interface{} {
	m := make(map[interface{}][]interface{})
	for _, x := range s.list {
		m[x] = nil
	}

	r := make([]interface{}, 0, 0)
	for k := range m {
		r = append(r, k)
	}
	return r
}

func (s *stream) GroupByInt(fn func(each interface{}) int64, r interface{}) {
	m := make(map[int64][]interface{})

	for _, x := range s.list {
		key := fn(x)
		if l, ok := m[key]; ok {
			l = append(l, x)
			m[key] = l
		} else {
			l = make([]interface{}, 0, 0)
			l = append(l, x)
			m[key] = l
		}
	}
	bytes, _ := json.Marshal(m)
	_ = json.Unmarshal(bytes, &r)
}

func (s *stream) GroupByString(fn func(each interface{}) string, r interface{}) {
	m := make(map[string][]interface{})

	for _, x := range s.list {
		key := fn(x)
		if l, ok := m[key]; ok {
			l = append(l, x)
			m[key] = l
		} else {
			l = make([]interface{}, 0, 0)
			l = append(l, x)
			m[key] = l
		}
	}
	bytes, _ := json.Marshal(m)
	_ = json.Unmarshal(bytes, &r)

}

func (s *stream) Sum(fn func(each interface{}) interface{}) float64 {
	var r float64 = 0
	for _, x := range s.list {
		p := fn(x)
		switch p.(type) {
		case string:
			break
		case int:
			r = r + (float64)(p.(int))
			break
		case float64:
			r = r + p.(float64)
			break
		}
	}
	return r
}

func (s *stream) Average(fn func(each interface{}) interface{}) float64 {
	var r float64 = 0
	for _, x := range s.list {
		p := fn(x)
		switch p.(type) {
		case string:
			break
		case int:
			r = r + (float64)(p.(int))
			break
		case float64:
			r = r + p.(float64)
			break
		}
	}
	return r / float64(len(s.list))
}

func (s *stream) Max(fn func(each interface{}) interface{}) float64 {
	var r float64 = math.MinInt64
	for _, x := range s.list {
		p := fn(x)
		switch p.(type) {
		case string:
			break
		case int:
			r = math.Max(r, (float64)(p.(int)))
			break
		case float64:
			r = math.Max(r, p.(float64))
			break
		}
	}
	return r
}

func (s *stream) Min(fn func(each interface{}) interface{}) float64 {
	var r = math.MaxFloat64
	for _, x := range s.list {
		p := fn(x)
		switch p.(type) {
		case string:
			break
		case int:
			r = math.Min(r, (float64)(p.(int)))
			break
		case float64:
			r = math.Min(r, p.(float64))
			break
		}
	}
	return r
}

func (s *stream) Reduce(initialValue interface{}, fn func(pre interface{}, cur interface{}) interface{}) interface{} {
	for i := 0; i < len(s.list); i++ {
		initialValue = fn(initialValue, s.list[i])
	}
	return initialValue
}
