package leetcode107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := dfs(root, make([][]int, 0), 0)

	length := len(res)

	for i := length/2 - 1; i >= 0; i-- {
		opp := length - 1 - i
		res[i], res[opp] = res[opp], res[i]
	}

	return res
}

func dfs(root *TreeNode, res [][]int, level int) [][]int {
	if root == nil {
		return res
	}

	if len(res)-1 < level {
		res = append(res, append(make([]int, 0), root.Val))
	} else {
		res[level] = append(res[level], root.Val)
	}

	level++

	res = dfs(root.Left, res, level)
	res = dfs(root.Right, res, level)
	return res
}
