package build_tree

// 剑指Offer 07 重建二叉树
// 输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	head := new(TreeNode)
	head.Val = preorder[0]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == head.Val {
			break
		}
	}

	head.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	head.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return head
}
