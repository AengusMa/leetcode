package leetcode111

// TreeNode 节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return helper(root, 0, sum)
}

func helper(node *TreeNode, cur int, dest int) bool {
	if node.Left != nil {
		if helper(node.Left, cur+node.Val, dest) {
			return true
		}
	}
	if node.Right != nil {
		if helper(node.Right, cur+node.Val, dest) {
			return true
		}
	}
	if node.Left == nil && node.Right == nil {
		return (cur+node.Val == dest)
	}
	return false
}
