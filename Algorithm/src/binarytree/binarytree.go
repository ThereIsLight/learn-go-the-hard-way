package main

import "fmt"

/*
二叉树是每个结点最多有两个子树的树结构。通常子树被称作“左子树”（left subtree）和“右子树”（right subtree）。二叉树常被用于实现二叉查找树和二叉堆。
*/
type Node struct {
	Val int  //就是小写会怎么样？？
	left, right *Node
}
func CreateTree(values []int) []Node {
	d := make([]Node, 0)  // 这是创建了一个切片吧
	for _, value := range values {
		d = append(d, Node{Val:value})
	}
	for i:=0; i<len(values)/2; i++ {  // /2??
		d[i].left = &d[i*2+1]
		if i*2+2 < len(values) {  //为什么i*2+2会越界呢？
			d[i].right = &d[i*2+2]
		}
	}
	return d
}
func (node *Node)Print() {
	fmt.Println(node.Val, " ")
}

// 前序遍历
func (node *Node) PreOrder() {
	if node == nil {
		return
	}
	node.Print()
	node.left.PreOrder()
	node.right.PreOrder()
}
//中序遍历
func (node *Node) MiddleOrder() {
	if node == nil {
		return
	}
	node.left.MiddleOrder()
	node.Print()
	node.right.MiddleOrder()
}
//后续遍历
func (node *Node) PostOrder() {
	if node == nil {
		return
	}
	node.left.PostOrder()
	node.right.PostOrder()
	node.Print()
}


func main() {
	values := []int{1,2,2,3,4,4,3}
	fmt.Println(values)
	t := CreateTree(values)
	fmt.Println(t)
	t[0].MiddleOrder()
}