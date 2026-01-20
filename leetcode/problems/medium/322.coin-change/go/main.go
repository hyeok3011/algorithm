package main

// https://leetcode.com/problems/coin-change/?envType=problem-list-v2&envId=rab78cw1
// 다른 사람의 DP풀이를 참고하여 다시 풀어본 DP방식
func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }

    sort.Ints(coins)
    dp := make([]int, amount + 1)
    for i := 1; i < len(dp); i++ {
        dp[i] = amount + 1
    }
    for _, coin := range coins {
        for i := coin; i < len(dp); i++ {
            dp[i] = min(dp[i], dp[i - coin] + 1)
        }
    }
  if dp[amount] == amount + 1 {
        return -1
    }
    return dp[amount]
}

func coinChangeBfs(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    
    sort.Ints(coins)

    queue := []int{amount}
    visited := make([]bool, amount+1)
    visited[amount] = true
    depth := 0
    for len(queue) > 0 {
        depth++
        n := len(queue)
        
        for i := 0; i < n; i++ {
            money := queue[i]
            
            for _, coin := range coins {
                newMoney := money - coin
                if newMoney == 0 {
                    return depth
                }

                if newMoney < 0 {
                    break 
                }
                
                if !visited[newMoney] {
                    visited[newMoney] = true
                    queue = append(queue, newMoney)
                }
            }
        }
        queue = queue[n:]
    }
    
    return -1
}
