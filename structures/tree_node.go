package structures

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTreeFromArray(arr []any) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) && len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(arr) && arr[i] != nil {
			left := &TreeNode{Val: arr[i].(int)}
			curr.Left = left
			queue = append(queue, left)
		}
		i++

		// Right child
		if i < len(arr) && arr[i] != nil {
			right := &TreeNode{Val: arr[i].(int)}
			curr.Right = right
			queue = append(queue, right)
		}
		i++
	}

	return root
}
