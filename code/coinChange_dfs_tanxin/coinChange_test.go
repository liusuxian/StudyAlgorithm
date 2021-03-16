package coinChange_dfs_tanxin

import (
    "math"
    "sort"
    "testing"
)

func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }

    ret := math.MaxInt64
    sort.Slice(coins, func(i, j int) bool {
        return coins[i] > coins[j]
    })
    dfs(coins, amount, 0, 0, &ret)

    if ret == math.MaxInt64 {
        return -1
    }

    return ret
}

func dfs(coins []int, amount, count, index int, ret *int) {
    if amount == 0 {
        *ret = min(*ret, count)
        return
    }

    if index == len(coins) {
        return
    }

    for k := amount / coins[index]; k >= 0 && k+count < *ret; k-- {
        dfs(coins, amount-k*coins[index], k+count, index+1, ret)
    }
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
    t.Log(coinChange([]int{1, 2, 5}, 100))
    t.Log(coinChange([]int{186, 419, 83, 408}, 6249))
}
