package leetcode111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var rst [][]int
	if root != nil {
		if root.Right == nil && root.Left == nil {
			if sum == root.Val {
				rst = append(rst, []int{root.Val})
			}
		} else {
			var tmp [][]int
			tmp = pathSum(root.Right, sum-root.Val)
			if len(tmp) > 0 {
				for _, v := range tmp {
					rst = append(rst, append([]int{root.Val}, v...))
				}
			}
			tmp = pathSum(root.Left, sum-root.Val)
			if len(tmp) > 0 {
				for _, v := range tmp {
					rst = append(rst, append([]int{root.Val}, v...))
				}
			}
		}
	}
	return rst
}
