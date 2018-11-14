package main

import (
	"fmt"
	"time"
)
func sample(message chan string) {
	// fmt.Println("hello goroutine !!!")  // 协程里面无法用到标准输入输出
	message <- "hello goroutine111 !!!"
	message <- "hello goroutine222 !!!"
	message <- "hello goroutine333 !!!"
	message <- "hello goroutine444 !!!"

}
func sample2(message chan string) {
	time.Sleep(2 * time.Second)
	str := <- message
	str += "I'm goroutine"
	message <- str
	close(message)
}
func main() {
	var message = make(chan string , 3)
	go sample(message)
	go sample2(message)
	time.Sleep(time.Second * 3)
	for str := range message {  //一直等到chan被关闭
		fmt.Println(str)
	}
	// close(message)
	fmt.Println("hello world !!!")
}
// select在程序内有多个队列的情况下做选择用