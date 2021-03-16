package minPathSum_dfs_memory

import (
    "math"
    "testing"
)

// 最小路径和 DFS + 记忆化搜索
func minPathSum(grid [][]int) int {
    row, col := len(grid), len(grid[0])
    memo := make([][]int, row)
    for k := range memo {
        memo[k] = make([]int, col)
        for i := 0; i < col; i++ {
            memo[k][i] = -1
        }
    }

    return dfs(grid, &memo, 0, 0)
}

func dfs(grid [][]int, memo *[][]int, row, col int) int {
    // 若越界，则认为不可达，距离为无穷大
    if !inArea(grid, row, col) {
        return math.MaxInt64
    }
    if (*memo)[row][col] > -1 {
        return (*memo)[row][col]
    }
    // 若到达终点，终点的贡献值是其本身
    if row == len(grid)-1 && col == len(grid[0])-1 {
        return grid[row][col]
    }
    // 取两者的较小值，计算出当前点的最小路径值
    (*memo)[row][col] = min(dfs(grid, memo, row, col+1), dfs(grid, memo, row+1, col)) + grid[row][col]
    return (*memo)[row][col]
}

// 判断坐标 (r, c) 是否在网格中
func inArea(grid [][]int, row, col int) bool {
    if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
        return false
    }

    return true
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}

func TestFun(t *testing.T) {
    t.Log(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
    t.Log(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
