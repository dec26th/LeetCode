package is_symmetric

// 剑指 Offer 28. 对称的二叉树

// 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return judgeSymmetric(root.Left, root.Right)
}

func judgeSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil {
		return false
	}
	if right == nil && left != nil {
		return false
	}

	if left.Val == right.Val {
		if judgeSymmetric(left.Left, right.Right) && judgeSymmetric(left.Right, right.Left) {
			return true
		}
	}
	return false
}
