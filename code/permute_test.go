package code

import "testing"

func permute(nums []int) [][]int {
    ret := make([][]int, 0)
    used := make([]bool, len(nums))

    dfs(nums, []int{}, used, &ret)
    return ret
}

func dfs(nums, path []int, used []bool, ret *[][]int) {
    if len(path) == len(nums) {
        tempPath := make([]int, len(path))
        copy(tempPath, path)
        *ret = append(*ret, tempPath)
        return
    }

    for i := 0; i < len(nums); i++ {
        if !used[i] {
            path = append(path, nums[i])
            used[i] = true
            dfs(nums, path, used, ret)
            // 回溯的过程中，将当前的节点从 path 中删除
            path = path[:len(path)-1]
            used[i] = false
        }
    }
}

func TestPermute(t *testing.T) {
    for _, v := range permute([]int{1, 2, 3}) {
        t.Log(v)
    }
}
