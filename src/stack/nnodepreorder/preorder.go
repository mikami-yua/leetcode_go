package nnodepreorder

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	queue := make([]*Node, 0)
	help(root, &queue)
	font := 0
	ans := make([]int, 0)
	for len(queue) > 0 {
		tempNode := queue[font]
		queue = queue[font+1:] //队列的缩减方式，append入队
		ans = append(ans, tempNode.Val)
	}
	return ans
}
func help(tempNode *Node, queue *[]*Node) {
	if tempNode == nil {
		return
	}
	*queue = append(*queue, tempNode)
	for i := 0; i < len(tempNode.Children); i++ {
		help(tempNode.Children[i], queue)
	}
}
