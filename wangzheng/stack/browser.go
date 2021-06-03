package stack

import (
	"fmt"
)

type Browser struct {
	Url          interface{}
	BackStack    *LinkStack
	ForwardStack *LinkStack
}

func NewLinkBrowser() *Browser {
	return &Browser{
		BackStack:    &LinkStack{},
		ForwardStack: &LinkStack{},
	}
}

func (b *Browser) CanForward() bool {
	if b.ForwardStack.IsEmpty() {
		return false
	}
	return true
}

func (b *Browser) CanBack() bool {
	if b.BackStack.IsEmpty() {
		return false
	}
	return true
}

func (b *Browser) Open(url string) {
	fmt.Printf("open url: %s \n", url)
	if b.Url != nil {
		b.BackStack.Push(b.Url)
	}
	b.Url = url
	b.ForwardStack.Flush()
}

func (b *Browser) Back() {
	b.ForwardStack.Push(b.Url)
	b.Url = b.BackStack.Pop()
	fmt.Println("后退到：", b.Url)

}

func (b *Browser) Forward() {
	b.BackStack.Push(b.Url)
	b.Url = b.ForwardStack.Pop()
	fmt.Println("前进到：", b.Url)
}
