# learning_algorithm
learning algorithm

##### 网格 DFS 遍历的模版代码：
``` go
func dfs(grid [][]byte, r, c int) {
    // 判断 base case
    // 如果坐标 (r, c) 超出了网格范围，直接返回
    if !inArea(grid, r, c) {
        return
    }
    // 访问上、下、左、右四个相邻结点
    dfs(grid, r-1, c)
    dfs(grid, r+1, c)
    dfs(grid, r, c-1)
    dfs(grid, r, c+1)
}

// 判断坐标 (r, c) 是否在网格中
func inArea(grid [][]byte, r, c int) bool {
    if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
        return false
    }
    
    return true
}
```
##### 递归代码模版
``` go
func recur(level int, param int) {
    // terminator
    if level > MAX_LEVEL {
        // process result
        return
    }

    // process current logic
    process(level, param)

    // drill down
    recur(level+1, newParam)

    // restore current status
}
```
##### 分治代码模版
``` go
func divide_conquer(problem *Problem, param int) int {
    // recursion terminator
    if problem == nil {
        process_result
        return return_result
    }

    // process current problem
    subproblems := split_problem(problem, data)
    subresult1 := divide_conquer(subproblem[0], p1)
    subresult2 := divide_conquer(subproblem[1], p1)
    subresult3 := divide_conquer(subproblem[2], p1)
    ...

    // merge
    result := process_result(subresult1, subresult2, subresult3)
    // revert the current level status

    return result
}
```
##### 二分查找代码模版
``` go
func binarySearch(slice []int, target int) int {
    left, right := 0, len(slice)-1

    for left <= right {
        mid := (left + right) >> 1

        if slice[mid] == target {
            return mid
        } else if slice[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return -1
}
```
##### 回溯代码模版
``` go
func backtrack(nums []int) [][]int {
    ret := make([][]int, 0)

    dfs(nums, []int{}, &ret)
    return ret
}

func dfs(nums, path []int, ret *[][]int) {
    // recursion terminator
    if len(path) == len(nums) {
        // process_result
        tempPath := make([]int, len(path))
        copy(tempPath, path)
        *ret = append(*ret, tempPath)
        return
    }

    for i := 0; i < len(nums); i++ {
        path = append(path, nums[i])
        dfs(nums, path, used, ret)
        // 回溯的过程中，将当前的节点从 path 中删除
        path = path[:len(path)-1]
    }
}
```
##### 递归，回溯和DFS的区别
- 递归是一种算法结构，回溯是一种算法思想。
- 一个递归就是在函数中调用函数本身来解决问题。
- 回溯就是通过不同的尝试来生成问题的解，有点类似于穷举，但是和穷举不同的是回溯会"剪枝"，意思就是对已经知道错误的结果没必要再枚举接下来的答案了，
  比如一个有序数列1,2,3,4,5，我要找和为5的所有集合，从前往后搜索我选了1，然后2，然后选3 的时候发现和已经大于预期，那么4,5肯定也不行，这就是一种对搜索过程的优化。
- 回溯搜索是深度优先搜索（DFS）的一种。
- 对于某一个搜索树来说（搜索树是起记录路径和状态判断的作用），回溯和DFS，其主要的区别是，回溯法在求解过程中不保留完整的树结构，而深度优先搜索则记下完整的搜索树。
- 为了减少存储空间，在深度优先搜索中，用标志的方法记录访问过的状态，这种处理方法使得深度优先搜索法与回溯法没什么区别了。