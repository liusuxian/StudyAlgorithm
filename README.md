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
#####