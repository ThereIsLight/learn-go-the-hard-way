package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}
	if (p==nil && q!=nil) || (q==nil && p!=nil) {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	// 在这种情况下，只有pq不为nil，且p.Val == q.Val这一种情况了。
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)  // 这种写法真的是666
}
