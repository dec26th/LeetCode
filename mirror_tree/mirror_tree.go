package mirror_tree

// 剑指 Offer 27. 二叉树的镜像

//请完成一个函数，输入一个二叉树，该函数输出它的镜像。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	reverse(root)
	return root
}

func reverse(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	reverse(root.Left)
	reverse(root.Right)
}
