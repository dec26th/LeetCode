package is_sub_structure

import "testing"

func TestIsSubStructure(t *testing.T) {
	test := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   11,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	sub := &TreeNode{
		Val:  3,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	isSubStructure(test, sub)
}
