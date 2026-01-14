package main

// https://leetcode.com/problems/partition-equal-subset-sum/?envType=problem-list-v2&envId=rab78cw1
// @@@@ AI도움을 받아 성능 개선했으나 다시 풀어볼만함
func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {

		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}


func canPartition2(nums []int) bool {
    totalSum := 0
    for _, n := range nums {
        totalSum += n
    }

    if totalSum%2 != 0 {
        return false
    }
    
    targetSum := totalSum / 2
    
    reachable := make(map[int]bool)
    reachable[0] = true

    for _, num := range nums {
        newSums := []int{}
        
        for k := range reachable {
            sum := k + num
            if sum == targetSum {
                return true
            }
            if sum < targetSum {
                newSums = append(newSums, sum)
            }
        }
        
        for _, s := range newSums {
            reachable[s] = true
        }
    }

    return reachable[targetSum]
}
