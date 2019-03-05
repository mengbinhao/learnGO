package test

import "testing"

func TestFirstTry(t *testing.T) {
	t.Log("my first try")
}

func TestExchange(t *testing.T) {
	var a int = 1
	var b int = 2
	a,b = b,a
	t.Log(a, b)
}