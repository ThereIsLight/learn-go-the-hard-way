package main

import (
	"algorithm/qsort"
)

func  main()  {
	var v = []int{4,1,7,6,9,2,8,0,3,5}
	qsort.QuickSort(v)
	//bubblesort.BubbleSort(v)
}

/*  调用其他包的函数
Go中如果函数名的首字母大写，表示该函数是公有的，可以被其他程序调用；
如果首字母小写，该函数就是是私有的。

    使用自定义包中的函数
包名+函数名

    将项目目录设置为GoPath，这样可以使用自定义的包（可以找到自定义的包）？？？？？？？？？？？？？？？？？？？？？？？？？？？？？？？
 */