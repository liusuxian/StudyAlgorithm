package coinChange_backtrack

import (
    "math"
    "sort"
    "testing"
)

// 零钱兑换 回溯法 超时
func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }

    ret := math.MaxInt64
    sort.Slice(coins, func(i, j int) bool {
        return coins[i] > coins[j]
    })
    dfs(coins, amount, 0, &ret)

    if ret == math.MaxInt64 {
        return -1
    }

    return ret
}

func dfs(coins []int, amount, count int, ret *int) {
    if amount == 0 {
        *ret = min(*ret, count)
        return
    }

    for _, coin := range coins {
        if amount < coin {
            continue
        }

        amount -= coin
        dfs(coins, amount, count+1, ret)
        amount += coin
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
    // 以下超时
    //t.Log(coinChange([]int{1, 2, 5}, 100))
    //t.Log(coinChange([]int{186, 419, 83, 408}, 6249))
}
