package leetcode105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(pre []int, in []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	rootVal := pre[0]
	if len(pre) == 1 {
		var root TreeNode
		root.Val = rootVal
		return &root
	}
	var leftTree *TreeNode
	i := 0
	for ; i < len(in); i++ {
		if in[i] == rootVal {
			break
		}
	}
	if i != 0 {
		leftTree = buildTree(pre[1:i+1], in[0:i])
	}
	var rightTree *TreeNode
	if i+1 < len(pre) && i+1 < len(in) {
		rightTree = buildTree(pre[i+1:], in[i+1:])
	}
	var rootTree TreeNode
	rootTree.Val = rootVal
	rootTree.Left = leftTree
	rootTree.Right = rightTree
	return &rootTree

}
