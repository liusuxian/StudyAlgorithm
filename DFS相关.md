***
#### 题目
##### 77. 组合
#### 地址
##### https://leetcode-cn.com/problems/combinations/
#### 方法一：回溯
##### 复杂度分析
###### 时间复杂度：O(k^n*k)。
###### 空间复杂度：O(n)。
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
###### 时间复杂度：O(n*n!)。
###### 空间复杂度：O(n)。
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
###### 时间复杂度：O(n*n!)。
###### 空间复杂度：O(n)。
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