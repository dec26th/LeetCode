package is_sub_structure

// 剑指 Offer 26. 树的子结构

//输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return judgeSubTree(A, B)
}

func judgeSubTree(A, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	if A.Val == B.Val && isSubTree(A, B) {
		return true

	}
	if judgeSubTree(A.Left, B) {
		return true
	}
	if judgeSubTree(A.Right, B) {
		return true
	}
	return false
}

func isSubTree(A, B *TreeNode) bool {
	if A == nil && B != nil {
		return false
	}
	if A != nil && B == nil {
		return true
	}
	if A == nil && B == nil {
		return true
	}

	if A.Val == B.Val && isSubTree(A.Left, B.Left) && isSubTree(A.Right, B.Right) {
		return true
	} else {
		return false
	}
}
