package stack

type Stack interface {
	Pop() interface{}
	Push(interface{})
	Flush()
	Top() interface{}
	IsEmpty() bool
}
