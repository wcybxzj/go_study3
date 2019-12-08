package __constant

import (
	"testing"
)

const (
	Monday = 1 + iota
	Tuesday
	Wednesday //3
)

const (
	Readable = 1 << iota //1
	Writeable //2
	Executeable //4
)


func TestConstantTry(t *testing.T)  {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T)  {
	a:=1
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executeable == Executeable)
}