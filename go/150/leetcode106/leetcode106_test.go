package leetcode106

func Index(xs []int, n int) int {
	for i, x := range xs {
		if x == n {
			return i
		}
	}
	return -1
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) > 0 && len(postorder) > 0 {
		node := &TreeNode{Val: postorder[len(postorder)-1]}
		idx := Index(inorder, node.Val)

		leftInorder := inorder[:idx]
		leftPostorder := postorder[:len(leftInorder)]
		node.Left = buildTree(leftInorder, leftPostorder)

		rightInorder := inorder[idx+1:]
		rightPostorder := postorder[len(leftPostorder) : len(postorder)-1]
		node.Right = buildTree(rightInorder, rightPostorder)

		return node
	}
	return nil
}
