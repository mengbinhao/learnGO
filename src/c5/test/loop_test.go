package test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i :=0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("default")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even_TestSwitchCaseCondition")
		case i%2 == 1:
			t.Log("Odd_TestSwitchCaseCondition")
		default:
			t.Log("unknow")
		}
	}
}
