package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
	"time"
	"algorithm/qsort"
	"algorithm/bubblesort"
)

/*
主程序
- 获取并解析命令行输入；
- 从对应文件中读取输入数据；
- 调用对应的排序函数；
- 将排序的结果输出到对应的文件中；
- 打印排序所花费时间的信息。
*/

//flag包：快速解析命令行参数 https://studygolang.com/articles/3365

var infile *string = flag.String("i", "infile", "File contains values for sorting")  //*string是什么？？ 输出字符串的相率更高吗？
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort","Sort Algorithm")

//从文件中逐行读取数字，返回一个字符串切片
func readValues(infile string)(values []int, err error){
	file, err := os.Open(infile)  // infile为文件名
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)  //??

	values = make([]int, 0)  //初始元素个数为0  //没有返回值？？？

	for {
		line, isPrefix, err1 := br.ReadLine()  //返回值 line []byte, isPrefix bool, err error
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break  // err=io.EOF,文件全部读取完
		}
		if isPrefix {  //????这个标志符做什么的？？
			fmt.Println("A too long line, seems unexpected")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)  //返回值int, err 将字符串转化为数字;Itoa将数字转化为字符串

		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return // ???
}
//将文件
func writeValues(values []int, outfile string) error{
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the ouput file", outfile)
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
func main(){
	fmt.Println(*infile)
	flag.Parse()  // 解析命令行？？？
	fmt.Println(*infile)
	// 本来全局变量infile的值为infile（默认值），执行了flag.Parse()之后，变成了unsorted.dat

	// 与命令行相关的操作
	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algotithm = ", *algorithm)  //
	}
	//调用函数，从文本文件中逐行读取[]byte,转化为string，再转化为数字。
	values, err := readValues(*infile)  //*string 到底是什么鬼？？???
	if err == nil {
		//fmt.Println("Read Values", values)
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")

		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
// 怎么查库函数的源代码，或者是功能解释。
// *string 是什么类型的变量?? 为什么前面再加上一个*就代表了string
//    *point就是代表地址指向的字符串
// type 关键字可以在windows上面打开文本文件

//go build的对象是谁？？b包，还是go文件
//go build与go install的区别
//go test是如何使用的，默认name_test.go文件吗？
//如何读取其他文件夹下面的数据
//当我使用go build sorter时，到底发生了什么其他的文件夹下面的进行了什么操作。？？？？