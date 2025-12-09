package main

// https://leetcode.com/problems/path-sum-iii/description/?envType=study-plan-v2&envId=leetcode-75
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	nums := make([]int, 1000)
	path := 0
	var recursion func(node *TreeNode, dpeth int)
	recursion = func(node *TreeNode, depth int) {
		if node.Val == targetSum {
			path++
		}
		nums[depth] = node.Val
		for i := 0; i < depth; i++ {
			nums[i] += node.Val
			if nums[i] == targetSum {
				path++
			}
		}

		if node.Left != nil {
			recursion(node.Left, depth+1)
		}

		if node.Right != nil {
			recursion(node.Right, depth+1)
		}

		for i := 0; i < depth; i++ {
			nums[i] -= node.Val
		}
	}

	recursion(root, 0)
	return path
}

// 다른사람은 prefix를 사용하여 푸는것을 발견 @@@@@@@
func pathSumPrefixSum(root *TreeNode, targetSum int) int {
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	path := 0

	var recursion func(node *TreeNode, sum int)
	recursion = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum += node.Val

		if count, ok := prefixSum[sum-targetSum]; ok {
			path += count
		}

		prefixSum[sum]++

		recursion(node.Left, sum)
		recursion(node.Right, sum)

		prefixSum[sum]--
	}

	recursion(root, 0)
	return path
}
