# learning_algorithm
learning algorithm

##### 网格 DFS 遍历的模版代码：
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
        return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
    }
##### 递归代码模版
    func recur(level int, param int) {
        // terminator
        if level > MAX_LEVEL {
            // process result
            return
        }
        
        // process current logic
        process(level, param)
        
        // drill down
        recur(level + 1, newParam)
        
        // restore current status
    }
##### 分治代码模版
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
##### 二分查找代码模版
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