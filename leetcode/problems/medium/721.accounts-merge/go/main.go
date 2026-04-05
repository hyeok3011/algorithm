package main

import "sort"

// https://leetcode.com/problems/accounts-merge/description/
// @@@@@@@ rank까지 도입해서 다시 풀어보면 좋을듯함.
func find(parent []int, v int) int {
	if parent[v] != v {
		parent[v] = find(parent, parent[v])
	}

	return parent[v]
}

func union(parent []int, src, dst int) {
	srcRoot, dstRoot := find(parent, src), find(parent, dst)
	parent[srcRoot] = dstRoot
}

func accountsMerge(accounts [][]string) [][]string {
	parent := make([]int, len(accounts))
	for i := range parent {
		parent[i] = i
	}
	emailToAccount := make(map[string]int)
	for i := range accounts {
		account := accounts[i]
		for j := 1; j < len(account); j++ {
			email := account[j]
			if rootIndex, exist := emailToAccount[email]; exist {
				union(parent, i, rootIndex)
			} else {
				emailToAccount[email] = i
			}
		}
	}

	groupEmails := make(map[int][]string)
	for email, idx := range emailToAccount {
		root := find(parent, idx)
		groupEmails[root] = append(groupEmails[root], email)
	}

	result := [][]string{}
	for root := range groupEmails {
		name := accounts[root][0]
		sort.Strings(groupEmails[root])
		result = append(result, append([]string{name}, groupEmails[root]...))
	}

	return result
}
