***
#### 题目
##### 127. 单词接龙
#### 地址
##### https://leetcode-cn.com/problems/word-ladder/
#### 方法一：单向广度优先搜索
##### 复杂度分析
- 时间复杂度：O(N + N * M * 26)。其中 N 为 wordList 的长度，M 为列表中单词的长度。
- 空间复杂度：O(N)。其中 N 为 wordList 的长度。
##### Golang实现
``` go
func ladderLength(beginWord string, endWord string, wordList []string) int {
    // 将 wordList 存储到 wordSet 集合中
    wordSet := map[string]bool{}
    for _, word := range wordList {
        wordSet[word] = true
    }
    // 检查 endWord 是否在 wordSet 集合中
    if _, ok := wordSet[endWord]; !ok {
        return 0
    }

    queue := []string{beginWord}
    visited := map[string]bool{beginWord: true}
    step := 1

    for len(queue) > 0 {
        for i := len(queue); i > 0; i-- {
            // 出队
            word := []byte(queue[0])
            queue = queue[1:]

            for j := 0; j < len(word); j++ {
                oldChar := word[j] // 临时保存
                for k := byte('a'); k <= 'z'; k++ {
                    word[j] = k
                    newWord := string(word)
                    // 已经访问过了，跳过
                    if _, ok := visited[newWord]; ok {
                        continue
                    }
                    if newWord == endWord {
                        // 如果和最后一个元素匹配，直接返回
                        return step + 1
                    }
                    // 如果新的字符串在 wordSet 集合中存在，将其添加到队列，并标记为已访问
                    if _, ok := wordSet[newWord]; ok {
                        queue = append(queue, newWord)
                        visited[newWord] = true
                    }
                }
                // 还原
                word[j] = oldChar
            }
        }
        step++
    }

    return 0
}
```
#### 方法二：双向广度优先搜索
##### 复杂度分析
- 时间复杂度：O(N + N * M * 26)。其中 N 为 wordList 的长度，M 为列表中单词的长度。
- 空间复杂度：O(N)。其中 N 为 wordList 的长度。
##### Golang实现
``` go
func ladderLength(beginWord string, endWord string, wordList []string) int {
    // 将 wordList 存储到 wordSet 集合中
    wordSet := map[string]bool{}
    for _, word := range wordList {
        wordSet[word] = true
    }
    // 检查 endWord 是否在 wordSet 集合中
    if _, ok := wordSet[endWord]; !ok {
        return 0
    }

    // 从两端 BFS 遍历要用的队列
    beginQueue := []string{beginWord}
    endQueue := []string{endWord}
    // 两端已经遍历过的节点
    beginVisited := map[string]bool{beginWord: true}
    endVisited := map[string]bool{endWord: true}
    step := 1

    for len(beginQueue) > 0 && len(endQueue) > 0 {
        // 每次判断正向是否比负向的元素更多。我们选择元素更小的那个继续更新。
        if len(beginQueue) > len(endQueue) {
            beginQueue, endQueue = endQueue, beginQueue
            beginVisited, endVisited = endVisited, beginVisited
        }

        for i := len(beginQueue); i > 0; i-- {
            // 出队
            word := []byte(beginQueue[0])
            beginQueue = beginQueue[1:]

            for j := 0; j < len(word); j++ {
                oldChar := word[j] // 临时保存
                for k := byte('a'); k <= 'z'; k++ {
                    word[j] = k
                    newWord := string(word)
                    // 已经访问过了，跳过
                    if _, ok := beginVisited[newWord]; ok {
                        continue
                    }
                    // 两端遍历相遇，结束遍历，返回 step
                    if _, ok := endVisited[newWord]; ok {
                        return step + 1
                    }
                    // 如果新的字符串在 wordSet 集合中存在，将其添加到队列，并标记为已访问
                    if _, ok := wordSet[newWord]; ok {
                        beginQueue = append(beginQueue, newWord)
                        beginVisited[newWord] = true
                    }
                }
                // 还原
                word[j] = oldChar
            }
        }
        step++
    }

    return 0
}
```
***
#### 题目
##### 433. 最小基因变化
#### 地址
##### https://leetcode-cn.com/problems/minimum-genetic-mutation/
#### 方法一：单向广度优先搜索
##### 复杂度分析
- 时间复杂度：O(N + N * M * 4)。其中 N 为 bank 的长度，M 为列表中单词的长度。
- 空间复杂度：O(N)。其中 N 为 bank 的长度。
##### Golang实现
``` go
func minMutation(start string, end string, bank []string) int {
    // 将 bank 存储到 bankSet 集合中
    bankSet := map[string]bool{}
    for _, genetic := range bank {
        bankSet[genetic] = true
    }
    // 检查 end 是否在 bankSet 集合中
    if _, ok := bankSet[end]; !ok {
        return -1
    }
    // 检查 start 和 end 是否相同
    if start == end {
        return 0
    }

    dict := []byte{'A', 'C', 'G', 'T'}
    queue := []string{start}
    visited := map[string]bool{start: true}
    step := 0

    for len(queue) > 0 {
        for i := len(queue); i > 0; i-- {
            // 出队
            genetic := []byte(queue[0])
            queue = queue[1:]

            for j := 0; j < len(genetic); j++ {
                oldChar := genetic[j] // 临时保存
                for _, dv := range dict {
                    genetic[j] = dv // 修改为dict中的字符
                    newGenetic := string(genetic)
                    // 已经访问过了，跳过
                    if _, ok := visited[newGenetic]; ok {
                        continue
                    }
                    if newGenetic == end {
                        // 如果和最后一个元素匹配，直接返回
                        return step + 1
                    }
                    // 如果新的基因序列在 bankSet 集合中存在，将其添加到队列，并标记为已访问
                    if _, ok := bankSet[newGenetic]; ok {
                        queue = append(queue, newGenetic)
                        visited[newGenetic] = true
                    }
                }
                // 还原
                genetic[j] = oldChar
            }
        }
        step++
    }

    return -1
}
```
#### 方法二：双向广度优先搜索
##### 复杂度分析
- 时间复杂度：O(N + N * M * 4)。其中 N 为 bank 的长度，M 为列表中单词的长度。
- 空间复杂度：O(N)。其中 N 为 bank 的长度。
##### Golang实现
``` go
func minMutation(start string, end string, bank []string) int {
    // 将 bank 存储到 bankSet 集合中
    bankSet := map[string]bool{}
    for _, genetic := range bank {
        bankSet[genetic] = true
    }
    // 检查 end 是否在 bankSet 集合中
    if _, ok := bankSet[end]; !ok {
        return -1
    }
    // 检查 start 和 end 是否相同
    if start == end {
        return 0
    }

    dict := []byte{'A', 'C', 'G', 'T'}
    // 从两端 BFS 遍历要用的队列
    startQueue := []string{start}
    endQueue := []string{end}
    // 两端已经遍历过的节点
    startVisited := map[string]bool{start: true}
    endVisited := map[string]bool{end: true}
    step := 0

    for len(startQueue) > 0 && len(endQueue) > 0 {
        // 每次判断正向是否比负向的元素更多。我们选择元素更小的那个继续更新。
        if len(startQueue) > len(endQueue) {
            startQueue, endQueue = endQueue, startQueue
            startVisited, endVisited = endVisited, startVisited
        }

        for i := len(startQueue); i > 0; i-- {
            // 出队
            genetic := []byte(startQueue[0])
            startQueue = startQueue[1:]

            for j := 0; j < len(genetic); j++ {
                oldChar := genetic[j] // 临时保存
                for _, dv := range dict {
                    genetic[j] = dv // 修改为dict中的字符
                    newGenetic := string(genetic)
                    // 已经访问过了，跳过
                    if _, ok := startVisited[newGenetic]; ok {
                        continue
                    }
                    // 两端遍历相遇，结束遍历，返回 step
                    if _, ok := endVisited[newGenetic]; ok {
                        return step + 1
                    }
                    // 如果新的基因序列在 bankSet 集合中存在，将其添加到队列，并标记为已访问
                    if _, ok := bankSet[newGenetic]; ok {
                        startQueue = append(startQueue, newGenetic)
                        startVisited[newGenetic] = true
                    }
                }
                // 还原
                genetic[j] = oldChar
            }
        }
        step++
    }

    return -1
}
```
***
#### 题目
##### 515. 在每个树行中找最大值
#### 地址
##### https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/#/description
#### 方法一：广度优先搜索
##### 复杂度分析
- 时间复杂度：O(n)。
- 空间复杂度：O(n)。
##### Golang实现
``` go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    ret := make([]int, 0)
    if root == nil {
        return ret
    }

    queue := []*TreeNode{root}
    for len(queue) > 0 {
        maxVal := math.MinInt64
        for i := len(queue); i > 0; i-- {
            head := queue[0]
            queue = queue[1:]
            maxVal = max(maxVal, head.Val)

            if head.Left != nil {
                queue = append(queue, head.Left)
            }
            if head.Right != nil {
                queue = append(queue, head.Right)
            }
        }
        ret = append(ret, maxVal)
    }

    return ret
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
```
***