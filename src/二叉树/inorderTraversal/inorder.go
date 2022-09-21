package inorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	} else {
		ans := make([]int, 0)
		help(root, &ans)
		return ans
	}

}
func help(root *TreeNode, ans *[]int) {
	if root.Left != nil {
		help(root.Left, ans)
	}
	*ans = append(*ans, root.Val)
	if root.Right != nil {
		help(root.Right, ans)
	}
}
