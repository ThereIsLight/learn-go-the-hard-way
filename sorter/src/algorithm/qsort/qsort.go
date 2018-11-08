package qsort

func quickSort(values []int, low, high int){
	if low>high {  //貌似也不会出现low>high的情况
		return
	}
	key := partition(values, low, high)
	quickSort(values, low, key-1)  //出现数组越界的问题怎么办？？1)对于go来说没有什么数组越界的问题;2)真的会出现数组越界吗？不会的
	quickSort(values, key+1, high)
}
//快速排序的切分
func partition(values []int, low, high int)int{
	//fmt.Println(low, high)
	last := high
	key:= values[high]  //选区数组的最后一个值为切分值。partition函数的目的就是为切分值找到合适的位置。保证切分值左边的数字小于等于它，右边的数字大于等于它。
	// 原理我很清楚，就是头尾指针并行处理，直到相遇。但是我不清楚如何在顺序的程序中实现并行的操作。
	for low<high {
		for low<high && values[low]<=key {
			low++
		}
		for low<high && values[high]>=key {
			high--
		}
		values[low], values[high] = values[high], values[low]
	}
	values[low], values[last] = values[last], values[low]
	//fmt.Println(values)
	return low  //正是切分值所在的下标
}
func QuickSort(values []int)  {
	quickSort(values, 0, len(values)-1)

}
// 算法的参考链接：左右指针法 https://blog.csdn.net/qq_36528114/article/details/78667034
// 相关的算法很多，但是只有这个我看的最明白。其他的博客理论看的明白，但是自己写不出他们的代码。
// 不知道是现在的状态的问题，还是字的基础实在是太差，我基本上看不懂，想不明白。

//for循环的条件
//无限循环 for{}