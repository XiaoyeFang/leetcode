package skiplist

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//
//const (
//	//最高层数
//	MAX_LEVEL = 16
//)
//
////跳表节点结构体
//type skipListNode struct {
//	//跳表保存的值
//	v interface{}
//	//用于排序的分值
//	score int
//	//层高
//	level int
//	//每层前进指针
//	forwards []*skipListNode
//}
//
////新建跳表节点
//func newSkipListNode(v interface{}, score, level int) *skipListNode {
//	return &skipListNode{v: v, score: score, forwards: make([]*skipListNode, level, level), level: level}
//}
//
////跳表结构体
//type SkipList struct {
//	//跳表头结点
//	head *skipListNode
//	//跳表当前层数
//	level int
//	//跳表长度
//	length int
//}
//
////实例化跳表对象
//func NewSkipList() *SkipList {
//	//头结点，便于操作
//	head := newSkipListNode(0, math.MinInt32, MAX_LEVEL)
//	return &SkipList{head, 1, 0}
//}
//
////获取跳表长度
//func (sl *SkipList) Length() int {
//	return sl.length
//}
//
////获取跳表层级
//func (sl *SkipList) Level() int {
//	return sl.level
//}
//
////插入节点到跳表中
//func (sl *SkipList) Insert(v interface{}, score int) int {
//	if nil == v {
//		return 1
//	}
//
//	//查找插入位置
//	cur := sl.head
//	//记录每层的路径
//	update := [MAX_LEVEL]*skipListNode{}
//	for i := MAX_LEVEL - 1; i >= 0; i-- {
//		for nil != cur.forwards[i] {
//			if cur.forwards[i].v == v {
//				return 2
//			}
//			if cur.forwards[i].score > score {
//				update[i] = cur
//				break
//			}
//			cur = cur.forwards[i]
//		}
//		if nil == cur.forwards[i] {
//			update[i] = cur
//		}
//	}
//
//	//通过随机算法获取该节点层数
//	level := 1
//	for i := 1; i < MAX_LEVEL; i++ {
//		if rand.Int31()%7 == 1 {
//			level++
//		}
//	}
//
//	//创建一个新的跳表节点
//	newNode := newSkipListNode(v, score, level)
//
//	//原有节点连接
//	for i := 0; i <= level-1; i++ {
//		next := update[i].forwards[i]
//		update[i].forwards[i] = newNode
//		newNode.forwards[i] = next
//	}
//
//	//如果当前节点的层数大于之前跳表的层数
//	//更新当前跳表层数
//	if level > sl.level {
//		sl.level = level
//	}
//
//	//更新跳表长度
//	sl.length++
//
//	return 0
//}
//
////查找
//func (sl *SkipList) Find(v interface{}, score int) *skipListNode {
//	if nil == v || sl.length == 0 {
//		return nil
//	}
//
//	cur := sl.head
//	for i := sl.level - 1; i >= 0; i-- {
//		for nil != cur.forwards[i] {
//			if cur.forwards[i].score == score && cur.forwards[i].v == v {
//				return cur.forwards[i]
//			} else if cur.forwards[i].score > score {
//				break
//			}
//			cur = cur.forwards[i]
//		}
//	}
//
//	return nil
//}
//
////删除节点
//func (sl *SkipList) Delete(v interface{}, score int) int {
//	if nil == v {
//		return 1
//	}
//
//	//查找前驱节点
//	cur := sl.head
//	//记录前驱路径
//	update := [MAX_LEVEL]*skipListNode{}
//	for i := sl.level - 1; i >= 0; i-- {
//		update[i] = sl.head
//		for nil != cur.forwards[i] {
//			if cur.forwards[i].score == score && cur.forwards[i].v == v {
//				update[i] = cur
//				break
//			}
//			cur = cur.forwards[i]
//		}
//	}
//
//	cur = update[0].forwards[0]
//	for i := cur.level - 1; i >= 0; i-- {
//		if update[i] == sl.head && cur.forwards[i] == nil {
//			sl.level = i
//		}
//
//		if nil == update[i].forwards[i] {
//			update[i].forwards[i] = nil
//		} else {
//			update[i].forwards[i] = update[i].forwards[i].forwards[i]
//		}
//	}
//
//	sl.length--
//
//	return 0
//}
//
//func (sl *SkipList) String() string {
//	return fmt.Sprintf("level:%+v, length:%+v", sl.level, sl.length)
//}

const UP_LEVELS_ABILITY = 16
const UP_LEVELS_TOTAL = 32

type skipListNode struct {
	score int64
	val   interface{}
	next  *skipListNode
	pre   *skipListNode
	up    *skipListNode
	down  *skipListNode
}

type skipList struct {
	head   *skipListNode
	tail   *skipListNode
	size   int
	levels int
}

func NewSkipList() *skipList {
	sl := new(skipList)
	sl.head = new(skipListNode)
	sl.tail = new(skipListNode)
	sl.head.score = math.MinInt64
	sl.tail.score = math.MaxInt64

	sl.head.next = sl.tail
	sl.tail.pre = sl.head

	sl.size = 0
	sl.levels = 1

	return sl
}

func (sl *skipList) Size() int {
	return sl.size
}

func (sl *skipList) Levels() int {
	return sl.levels
}

func (sl *skipList) Get(score int64) interface{} {
	node := sl.findNode(score)
	if node.score == score {
		return node.val
	} else {
		return nil
	}
}

func (sl *skipList) Insert(score int64, val interface{}) {
	f := sl.findNode(score)
	if f.score == score {
		f.val = val
		return
	}
	curNode := new(skipListNode)
	curNode.score = score
	curNode.val = val

	sl.insertAfter(f, curNode)

	rander := rand.New(rand.NewSource(time.Now().UnixNano()))

	curlevels := 1
	for rander.Intn(UP_LEVELS_TOTAL) < UP_LEVELS_ABILITY {
		curlevels++
		if curlevels > sl.levels {
			//每高一层，在上面就新建一层索引
			sl.newlevels()
		}
        // 每新建一层索引就向上查找当前插入节点的前驱
		for f.up == nil {
			f = f.pre
		}
		f = f.up
		tmpNode := &skipListNode{score: score,val: val}

		curNode.up = tmpNode
		tmpNode.down = curNode
		sl.insertAfter(f, tmpNode)

		curNode = tmpNode
	}

	sl.size++
}

func (sl *skipList) Remove(score int64) interface{} {
	f := sl.findNode(score)
	if f.score != score {
		return nil
	}
	v := f.val

	for f != nil {
		f.pre.next = f.next
		f.next.pre = f.pre
		f = f.up
	}
	return v
}

func (sl *skipList) newlevels() {
	nhead := &skipListNode{score: math.MinInt64}
	ntail := &skipListNode{score: math.MaxInt64}
	nhead.next = ntail
	ntail.pre = nhead

	sl.head.up = nhead
	nhead.down = sl.head
	sl.tail.up = ntail
	ntail.down = sl.tail

	sl.head = nhead
	sl.tail = ntail
	sl.levels++
}

func (sl *skipList) insertAfter(pNode *skipListNode, curNode *skipListNode) {
	curNode.next = pNode.next
	curNode.pre = pNode
	pNode.next.pre = curNode
	pNode.next = curNode
}

func (sl *skipList) findNode(score int64) *skipListNode {
	p := sl.head

	for p != nil {
		if p.score == score {
			if p.down == nil {
				return p
			}
			p = p.down
		} else if p.score < score {
			if p.next.score > score {
				if p.down == nil {
					return p
				}
				p = p.down
			} else {
				p = p.next
			}
		}
	}
	return p
}

func (sl *skipList) Print() {

	mapScore := make(map[int64]int)

	p := sl.head
	for p.down != nil {
		p = p.down
	}
	index := 0
	for p != nil {
		mapScore[p.score] = index
		p = p.next
		index++
	}
	p = sl.head
	for i := 0; i < sl.levels; i++ {
		q := p
		preIndex := 0
		for q != nil {
			s := q.score
			if s == math.MinInt64 {
				fmt.Printf("%s", "BEGIN")
				q = q.next
				continue
			}
			index := mapScore[s]
			c := (index - preIndex - 1) * 6
			for m := 0; m < c; m++ {
				fmt.Print("-")
			}
			if s == math.MaxInt64 {
				fmt.Printf("-->%s\n", "END")
			} else {
				fmt.Printf("-->%3d", s)
				preIndex = index
			}
			q = q.next
		}
		p = p.down
	}
}