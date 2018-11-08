package bubblesort

func BubbleSort(values []int){
	var flag = true  // 减少比较次数
	//fmt.Println(values)
	var n = len(values)
	var i,j int
	for i=0; i<n-1; i++ {
		for j=0; j<n-i-1; j++ {
			if values[j] > values[j+1] {
				values[j+1], values[j] = values[j], values[j+1]  //与Python一样
				flag = false  //
			}
		}
		if flag {
			break
		}
		flag = true
		//fmt.Println(values)
	}
	//fmt.Println(values)
}

//函数的返回值呢？？ 没有返回值
//传的值还是引用？？ 传的是值
//flag的作用:如果if判断一直不成立，则说明数组已经有序。不需要进行循环来比较。
