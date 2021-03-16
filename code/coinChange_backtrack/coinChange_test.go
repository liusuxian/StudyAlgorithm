package coinChange_backtrack

import (
    "math"
    "sort"
    "testing"
)

func coinChange(coins []int, amount int) int {
    ret := math.MaxInt64

    sort.Slice(coins, func(i, j int) bool {
        return i > j
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
        //if amount < coin {
        //    continue
        //}
        //
        //amount -= coin
        //dfs(coins, amount, count+1, ret)
        //amount += coin

        k := amount / coin
        if k <= 0 {
            continue
        }

        amount -= coin * k
        dfs(coins, amount, count+k, ret)
        amount += coin * k
    }
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}

func TestFun(t *testing.T) {
    //t.Log(coinChange([]int{1, 2, 5}, 11))
    //t.Log(coinChange([]int{2}, 3))
    //t.Log(coinChange([]int{1}, 0))
    //t.Log(coinChange([]int{1}, 1))
    //t.Log(coinChange([]int{1}, 2))
    //t.Log(coinChange([]int{1, 2, 5}, 100))
    t.Log(coinChange([]int{186, 419, 83, 408}, 6249))
}
