package leetcode110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, ok := getDepthAndBalance(root)
	return ok
}

func getDepthAndBalance(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	d1, ok1 := getDepthAndBalance(root.Left)
	d2, ok2 := getDepthAndBalance(root.Right)
	if !ok1 || !ok2 {
		return 0, false
	}
	if abs(d1-d2) > 1 {
		return 0, false
	}

	return max(d1+1, d2+1), true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
