package coinChange_dfs_memory

import (
    "fmt"
    "math"
    "testing"
)

func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }

    ret := math.MaxInt64
    dfs(coins, amount, 0, &ret)

    if ret == math.MaxInt64 {
        return -1
    }

    return ret
}

func dfs(coins []int, amount, count int, ret *int) {
    fmt.Println(amount, count, *ret)
    if amount < 0 {
        return
    }

    if amount == 0 {
        *ret = min(*ret, count)
        return
    }

    for _, coin := range coins {
        dfs(coins, amount-coin, count+1, ret)
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
    //t.Log(coinChange([]int{2}, 3))
    //t.Log(coinChange([]int{1}, 0))
    //t.Log(coinChange([]int{1}, 1))
    //t.Log(coinChange([]int{1}, 2))
    //t.Log(coinChange([]int{1, 2, 5}, 100))
    //t.Log(coinChange([]int{186, 419, 83, 408}, 6249))
}
