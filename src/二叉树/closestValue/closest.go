package closestValue

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
定义当前节点值val、最小差值时的节点值closest
如果当前差值比之前的最小差值还要小，更新closest
target小于当前值，遍历左子树，否则遍历右子树

*/

type pair struct {
	minDiff float64
	node    *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	ansPair := pair{minDiff: math.MaxFloat64, node: root}
	help(root, target, &ansPair)
	return ansPair.node.Val
}

func help(root *TreeNode, target float64, ansPair *pair) {
	diff := math.Abs(float64(root.Val) - target)
	if diff < ansPair.minDiff {
		ansPair.minDiff = diff
		ansPair.node = root
	}
	if target < float64(root.Val) && root.Left != nil {
		help(root.Left, target, ansPair)
	} else if target > float64(root.Val) && root.Right != nil {
		help(root.Right, target, ansPair)
	} else {
		return
	}
}
