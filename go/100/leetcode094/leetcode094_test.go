package leetcode094

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var xs []int
	if root != nil {
		xs = append(xs, inorderTraversal(root.Left)...)
		xs = append(xs, root.Val)
		xs = append(xs, inorderTraversal(root.Right)...)
	}
	return xs
}
