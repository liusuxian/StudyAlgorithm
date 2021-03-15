package code

import "testing"

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

func TestClimbStairs(t *testing.T) {
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
