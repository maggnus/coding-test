/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return depth(1, root)
}

func depth(d int, node *TreeNode) int {

	if node.Left == nil && node.Right == nil {
		return d
	}

	var dl, dr int

	if node.Left != nil {
		dl = depth(d+1, node.Left)
	}

	if node.Right != nil {
		dr = depth(d+1, node.Right)
	}

	if dl >= dr {
		return dl
	} else {
		return dr
	}
}