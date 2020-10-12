package lowest_common_ancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, q, p *TreeNode) *TreeNode {
	if root == nil || q == root || p == root {
		return root
	}
	left := lowestCommonAncestor(root.Left, q, p)
	right := lowestCommonAncestor(root.Right, q, p)
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}
	return root

}
