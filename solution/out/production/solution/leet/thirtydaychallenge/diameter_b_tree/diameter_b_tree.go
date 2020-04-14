package diameterOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxLen = 0

func diameterOfBinaryTree(root *TreeNode) int {
	maxLen = 0
	traverse(root)
	return maxLen
}

func traverse(root *TreeNode) int {
	if root == nil {
		return -1
	}

	left := 1 + traverse(root.Left)
	right := 1 + traverse(root.Right)

	if (left + right) > maxLen {
		maxLen = left + right
	}

	return getMax(left, right)
}

func getMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
