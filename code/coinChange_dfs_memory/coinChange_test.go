package coinChange_dfs_memory

import (
    "math"
    "sort"
    "testing"
)

func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }

    // memo[n] 表示钱币n可以被换取的最少的硬币数，不能换取就为-1
    memo := make([]int, amount)
    sort.Slice(coins, func(i, j int) bool {
        return coins[i] > coins[j]
    })
    return dfs(coins, memo, amount)
}

func dfs(coins []int, memo []int, amount int) int {
    if amount < 0 {
        return -1
    }

    if amount == 0 {
        return 0
    }

    if memo[amount-1] != 0 {
        return memo[amount-1]
    }

    minVal := math.MaxInt64
    for _, coin := range coins {
        ret := dfs(coins, memo, amount-coin)
        if ret >= 0 && ret < minVal {
            minVal = ret + 1
        }
    }

    if minVal == math.MaxInt64 {
        memo[amount-1] = -1
    } else {
        memo[amount-1] = minVal
    }

    return memo[amount-1]
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}

func TestFun(t *testing.T) {
    t.Log(coinChange([]int{1, 2, 5}, 11))
    t.Log(coinChange([]int{2}, 3))
    t.Log(coinChange([]int{1}, 0))
    t.Log(coinChange([]int{1}, 1))
    t.Log(coinChange([]int{1}, 2))
    // 以下超时
    t.Log(coinChange([]int{1, 2, 5}, 100))
    t.Log(coinChange([]int{186, 419, 83, 408}, 6249))
}
