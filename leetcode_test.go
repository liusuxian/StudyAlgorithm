package learning_algorithm

import (
    "math"
    "testing"
)

func climbStairs(n int) int {
    dp := make([]int, n+4)
    dp[1], dp[2], dp[3] = 1, 2, 4

    if n <= 3 {
        return dp[n]
    }

    for i := 4; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
    }

    return dp[n]
}

func climbStairs1(n int) int {
    dp := make([][4]int, n+4)
    dp[1][1] = 1
    dp[2][2] = 1
    dp[3][1] = 1
    dp[3][2] = 1
    dp[3][3] = 1

    for i := 4; i <= n; i++ {
        if i-1 >= 0 {
            dp[i][1] = dp[i-1][2] + dp[i-1][3]
        }
        if i-2 >= 0 {
            dp[i][2] = dp[i-2][1] + dp[i-2][3]
        }
        if i-3 >= 0 {
            dp[i][3] = dp[i-3][1] + dp[i-3][2]
        }
    }

    return dp[n][1] + dp[n][2] + dp[n][3]
}

func maxSumSubmatrix(matrix [][]int, k int) int {
    m, n, maxArea := len(matrix), len(matrix[0]), math.MinInt64
    // O(cols ^ 2 * rows)
    for l := 0; l < n; l++ { // 枚举左边界
        rowSum := make([]int, m) // 左边界改变才算区域的重新开始
        for r := l; r < n; r++ { // 枚举右边界
            for i := 0; i < m; i++ { // 按每一行累计到 dp
                rowSum[i] += matrix[i][r]
            }
            maxArea = max(maxArea, dpMax(rowSum, k))
            if maxArea == k {
                return k // 尽量提前
            }
        }
    }

    return maxArea
}

// 在数组 arr 中，求不超过 k 的最大值
func dpMax(arr []int, k int) int {
    rollSum := arr[0]
    rollMax := rollSum

    // O(rows)
    for i := 1; i < len(arr); i++ {
        if rollSum > 0 {
            rollSum += arr[i]
        } else {
            rollSum = arr[i]
        }

        rollMax = max(rollMax, rollSum)
    }
    if rollMax <= k {
        return rollMax
    }
    // O(rows ^ 2)
    maxArea := math.MinInt64
    for l := 0; l < len(arr); l++ {
        sum := 0
        for r := l; r < len(arr); r++ {
            sum += arr[r]
            if sum > maxArea && sum <= k {
                maxArea = sum
            }
            if maxArea == k {
                return k // 尽量提前
            }
        }
    }

    return maxArea
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}

func TestFun(t *testing.T) {
    t.Log(climbStairs(0))
    t.Log(climbStairs(1))
    t.Log(climbStairs(2))
    t.Log(climbStairs(3))
    t.Log(climbStairs(4))
    t.Log(climbStairs(5))

    t.Log(climbStairs1(0))
    t.Log(climbStairs1(1))
    t.Log(climbStairs1(2))
    t.Log(climbStairs1(3))
    t.Log(climbStairs1(4))
    t.Log(climbStairs1(5))
}
