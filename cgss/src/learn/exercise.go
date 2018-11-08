package main

import "fmt"

func main() {
	ss := []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(ss)
	for i, v := range ss {
		if v == 5 {
			//ss = append(ss[:i], ss[i+1:]...)
			ss = append(ss[:i-1], ss[:i+1]...)
		}
	}
	fmt.Println(ss)
}


