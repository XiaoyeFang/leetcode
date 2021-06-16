package skiplist

import (
	"testing"
)

func TestNewSkipList(t *testing.T) {

}

type student struct {
	Name string
}
type people struct {
	s *student
}

func TestSkipList_Insert(t *testing.T) {
	//list :=NewSkipList()
	//list.Insert(1,22)
	//list.Insert(2,23)
	sk := NewSkipList()
	sk.findNode(22)

	sk.Insert(100, "lala")
	sk.Insert(11, "sx")
	//sk.Insert(22, "11")
	//sk.Insert(3, "dd")
	//sk.Insert(80, "bb")
	//sk.Insert(77, "bb")
	//sk.Insert(6, "bb")
	//sk.Insert(88, "bb")
	//sk.Insert(33, "bb")
	//sk.Insert(44, "bb")

	//fmt.Println(sk.Get(22))
	//fmt.Println(sk.Get(55))
	//fmt.Println(sk.Remove(22))
	//fmt.Println(sk.Get(22))
	//fmt.Println(sk.Size())
	//fmt.Println(sk.Layout())
	//sk.Print()

}
