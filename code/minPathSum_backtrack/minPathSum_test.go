package minPathSum_backtrack

import (
    "math"
    "testing"
)

// 最小路径和 回溯法 超时
func minPathSum(grid [][]int) int {
    dirs := [][]int{{1, 0}, {0, 1}}
    ret := math.MaxInt64

    dfs(grid, dirs, 0, 0, grid[0][0], &ret)
    return ret
}

func dfs(grid, dirs [][]int, row, col, sum int, ret *int) {
    if row == len(grid)-1 && col == len(grid[0])-1 {
        *ret = min(*ret, sum)
        return
    }

    for _, dir := range dirs {
        nextRow := row + dir[0]
        nextCol := col + dir[1]
        if !inArea(grid, nextRow, nextCol) {
            continue
        }

        sum += grid[nextRow][nextCol]
        dfs(grid, dirs, nextRow, nextCol, sum, ret)
        // 回溯
        sum -= grid[nextRow][nextCol]
    }
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
