package bubblesort

import "testing"

func TestBubbleSort(t *testing.T) {
	values := []int{5,4,3,2,1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 ||
		values[4] !=5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 4 5")
	}
}
func TestBubbleSort2(t *testing.T) {
	values := []int{5,5,3,2,1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 ||
		values[4] !=5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 4 5")
	}
}
func TestBubbleSort3(t *testing.T) {
	values := []int{5}
	BubbleSort(values)
	if values[0] !=5 {
		t.Error("BubbleSort() failed. Got", values, "Excepted 5")
	}
}

//testing这个包是专门用来做测试的吗？？
//输入test自动补全了后面的函数名TestBubbleSort,好像直到你要去测试什么一样。
//有点像JUnit，战门做单元测试的。