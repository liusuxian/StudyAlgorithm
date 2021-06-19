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
##### https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
#### 方法一：循环遍历，条件满足就选择
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
#### 题目
##### 283. 移动零
#### 地址
##### https://leetcode-cn.com/problems/move-zeroes/
#### 方法一：循环遍历，条件满足就选择
##### 复杂度分析
- 时间复杂度：O(n)。
- 空间复杂度：O(1)。
##### Golang实现
``` go
func moveZeroes(nums []int) {
    n := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[n] = nums[i]
            n++
        }
    }

    for ; n < len(nums); n++ {
        nums[n] = 0
    }
}
```
***
#### 题目
##### 206. 反转链表
#### 地址
##### https://leetcode-cn.com/problems/reverse-linked-list/
#### 方法一：循环遍历
##### 复杂度分析
- 时间复杂度：O(n)。
- 空间复杂度：O(1)。
##### Golang实现
``` go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var last *ListNode
    // 要改每条边，所以需要访问链表
    for head != nil {
        nextHead := head.Next
        // 改一条边
        head.Next = last
        // last，head 向后移动一位
        last, head = head, nextHead
    }
    
    return last
}
```
***
#### 题目
##### 25. K 个一组翻转链表
#### 地址
##### https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
#### 方法一：循环遍历
##### 复杂度分析
- 时间复杂度：O(n)。
- 空间复杂度：O(1)。
##### Golang实现
``` go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    protect := &ListNode{Next: head}
    // 分组（找到每一组的开始、结尾），按组遍历
    // last = 上一组结尾
    last := protect
    for head != nil {
        end := getEnd(head, k)
        if end == nil {
            break
        }

        nextGroupHead := end.Next
        // 处理 head 到 end 之间的 k-1 条边的反转
        reverseList(head, end)
        // 上一组跟本组的新开始（旧end）建立联系
        last.Next = end
        // 本组的新结尾（head）跟下一组建立联系
        head.Next = nextGroupHead
        // 分组遍历
        last, head = head, nextGroupHead
    }

    return protect.Next
}

func getEnd(head *ListNode, k int) *ListNode {
    for head != nil {
        k--
        if k == 0 {
            break
        }
        head = head.Next
    }

    return head
}

// head 到 end 之间反过来
func reverseList(head *ListNode, end *ListNode) {
    if head == end {
        return
    }

    last := head
    head = head.Next
    // 要改每条边，所以需要访问链表
    for head != end {
        nextHead := head.Next
        // 改一条边
        head.Next = last
        // last，head 向后移动一位
        last, head = head, nextHead
    }
    end.Next = last
}
```
***