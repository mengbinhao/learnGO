package test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr[3]int //just declare, default value is zero
	arr1 := [4]int{1,2,3,4} // declare and init
	arr3 := [...]int{1,2,3,4} //不指定元素个数
	//c := [2][2]int{{1,2}, {3,4}} 多维
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	for i := 0; i< len(arr3); i++ {
		t.Log(arr3[i])
	}
}

func TestArrayTravel2(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	//for idx, e := range arr3 {
	//索引必须占位
	for _, e := range arr3 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	//a[1:2]
	//a[1:]
	//a[:3]
	//a[:]
	//a[1:len(a)]
	arr3_sec := arr3[:]
	t.Log(arr3_sec)
}