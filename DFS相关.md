***
#### 题目
##### 77. 组合
#### 地址
##### https://leetcode-cn.com/problems/combinations/
#### 方法一：回溯
##### 复杂度分析
- 时间复杂度：O(k^n*k)。
- 空间复杂度：O(n)。
##### Golang实现
``` go
func combine(n int, k int) [][]int {
    combs := make([][]int, 0)
    backtrack(1, n, k, []int{}, &combs)
    return combs
}

func backtrack(index, n, k int, comb []int, combs *[][]int) {
    if len(comb) == k {
        newComb := make([]int, k)
        copy(newComb, comb)
        *combs = append(*combs, newComb)
        return
    }

    // 剪枝，取到后面数量不够了就直接跳过，可以节省很多时间
    for i := index; i <= n && k-len(comb) <= n-i+1; i++ {
        comb = append(comb, i)
        backtrack(i+1, n, k, comb, combs)
        comb = comb[:len(comb)-1]
    }
}
```
***
#### 题目
##### 46. 全排列
#### 地址
##### https://leetcode-cn.com/problems/permutations/
#### 方法一：回溯
##### 复杂度分析
- 时间复杂度：O(n*n!)。
- 空间复杂度：O(n)。
##### Golang实现
``` go
func permute(nums []int) [][]int {
    ret := make([][]int, 0)
    if len(nums) == 0 {
        return ret
    }

    used := make([]bool, len(nums))
    path := make([]int, 0, len(nums))
    backtrack(nums, path, used, &ret)

    return ret
}

func backtrack(nums, path []int, used []bool, ret *[][]int) {
    if len(path) == len(nums) {
        newPath := make([]int, len(path))
        copy(newPath, path)
        *ret = append(*ret, newPath)
        return
    }

    for i := 0; i < len(nums); i++ {
        if !used[i] {
            path = append(path, nums[i])
            used[i] = true

            backtrack(nums, path, used, ret)

            used[i] = false
            path = path[:len(path)-1]
        }
    }
}
```
***
#### 题目
##### 47. 全排列 II
#### 地址
##### https://leetcode-cn.com/problems/permutations-ii/
#### 方法一：回溯
##### 复杂度分析
- 时间复杂度：O(n*n!)。
- 空间复杂度：O(n)。
##### Golang实现
``` go
func permuteUnique(nums []int) [][]int {
    ret := make([][]int, 0)
    if len(nums) == 0 {
        return ret
    }

    sort.Ints(nums)
    used := make([]bool, len(nums))
    path := make([]int, 0, len(nums))
    backtrack(nums, path, used, &ret)

    return ret
}

func backtrack(nums, path []int, used []bool, ret *[][]int) {
    if len(path) == len(nums) {
        newPath := make([]int, len(path))
        copy(newPath, path)
        *ret = append(*ret, newPath)
        return
    }

    for i := 0; i < len(nums); i++ {
        if used[i] {
            continue
        } else if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
            continue
        }

        path = append(path, nums[i])
        used[i] = true

        backtrack(nums, path, used, ret)

        used[i] = false
        path = path[:len(path)-1]
    }
}
```
***
#### 题目
##### 22. 括号生成
#### 地址
##### https://leetcode-cn.com/problems/generate-parentheses/#/description
#### 方法一：回溯（加法）
##### 复杂度分析
- 时间复杂度：O(4^n/sqrt(n))。
- 空间复杂度：O(n)。
##### Golang实现
``` go
func generateParenthesis(n int) []string {
    ret := make([]string, 0)
    backtrack("", 0, 0, n, &ret)
    return ret
}

func backtrack(curStr string, left, right, n int, ret *[]string) {
    if left == n && right == n {
        *ret = append(*ret, curStr)
        return
    }

    // 当 left 使用数量小于 right 使用数量时剪枝
    if left < right {
        return
    }

    if left < n {
        backtrack(curStr+"(", left+1, right, n, ret)
    }

    if right < n {
        backtrack(curStr+")", left, right+1, n, ret)
    }
}
```
#### 方法二：回溯（减法）
##### 复杂度分析
- 时间复杂度：O(4^n/sqrt(n))。
- 空间复杂度：O(n)。
##### Golang实现
``` go
func generateParenthesis(n int) []string {
    ret := make([]string, 0)
    backtrack("", n, n, &ret)
    return ret
}

func backtrack(curStr string, left, right int, ret *[]string) {
    if left == 0 && right == 0 {
        *ret = append(*ret, curStr)
        return
    }

    // 当 left 剩余数量大于 right 剩余数量时剪枝
    if left > right {
        return
    }

    if left > 0 {
        backtrack(curStr+"(", left-1, right, ret)
    }

    if right > 0 {
        backtrack(curStr+")", left, right-1, ret)
    }
}
```
***
#### 题目
##### 322. 零钱兑换
#### 地址
##### https://leetcode-cn.com/problems/coin-change/
#### 方法一：回溯
##### 复杂度分析
- 时间复杂度：O()。
- 空间复杂度：O()。
##### Golang实现
``` go
func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }

    sort.Ints(coins)
    ret := math.MaxInt64
    backtrack(&coins, amount, len(coins), 0, &ret)

    if ret == math.MaxInt64 {
        return -1
    }

    return ret
}

func backtrack(coins *[]int, amount, index, count int, ret *int) {
    if amount == 0 {
        *ret = min(*ret, count)
        return
    }
    if index == 0 {
        return
    }

    for i := amount / (*coins)[index-1]; i >= 0 && i+count < *ret; i-- {
        backtrack(coins, amount-i*(*coins)[index-1], index-1, i+count, ret)
    }
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}
```
***