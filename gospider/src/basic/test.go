package main

import "fmt"

func main() {
	var s = []int{1}
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)

}