package reflect

import (
	"fmt"
	"testing"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {
	for msg := range msgchan {
		m := msg.(int)
		fmt.Println("msg: ", m)
	}
}

func (p *Project) run(msgchan chan interface{}) {

	for {
		//defer p.deferError()
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}

func (p *Project) Main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("don't worry, I can take care of myself")
		}
	}()
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {

		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 10)
}

func Cantype(v interface{}) {

	if value, ok := v.(string); ok {
		fmt.Println(value)
	}
}

type Student struct {
	name string
}

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		fmt.Println("i",i)
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}


func TestReflectex(t *testing.T) {

	testReflect()

}
