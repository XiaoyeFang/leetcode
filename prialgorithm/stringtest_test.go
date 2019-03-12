package prialgorithm

import "testing"

func TestReverse(t *testing.T) {
	sum := Reverse(-4432)
	t.Log(sum)
}

func TestMyAtoi(t *testing.T) {
	str := "3.14159"
	num := MyAtoi(str)
	t.Log(num)
}
