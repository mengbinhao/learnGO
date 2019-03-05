package test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wenesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday, Wenesday)
}

func TestConst2(t *testing.T) {
	a := 7
	t.Log(a&Readable == Readable, a&Writeable  == Writeable, a&Executable == Executable)
}