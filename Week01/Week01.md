***
#### 题目
##### 88. 合并两个有序数组
#### 地址
##### https://leetcode-cn.com/problems/merge-sorted-array/
#### 方法一：双指针
##### 复杂度分析
- 时间复杂度：O(m+n)。
- 空间复杂度：O(1)。
##### Golang实现
``` go
func merge(nums1 []int, m int, nums2 []int, n int) {
    // 主体思路：i，j 两个指针倒着扫描，谁大选谁
    // 细节判断：i，j 不能越界（一个<0，就选另一个）
    i, j := m-1, n-1
    for k := m + n - 1; k >= 0; k-- {
        if j < 0 || (i >= 0 && nums1[i] > nums2[j]) {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
    }
}
```
***
#### 题目
##### 26. 删除有序数组中的重复项
#### 地址
##### https://leetcode-cn.com/problems/merge-sorted-array/
#### 方法一：循环遍历
##### 复杂度分析
- 时间复杂度：O(n)。
- 空间复杂度：O(1)。
##### Golang实现
``` go
func removeDuplicates(nums []int) int {
    n := 0
    for i := 0; i < len(nums); i++ {
        if i == 0 || nums[i] != nums[i-1] {
            nums[n] = nums[i]
            n++
        }
    }

    return n
}
```
***