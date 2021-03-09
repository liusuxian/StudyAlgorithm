***
#### 题目
##### 127. 单词接龙
#### 地址
##### https://leetcode-cn.com/problems/word-ladder/
#### 方法一：广度优先搜索 + 优化建图
##### 复杂度分析
###### 时间复杂度：O(N×C^2)。其中 N 为 wordList 的长度，C 为列表中单词的长度。
###### 空间复杂度：O(N×C^2)。其中 N 为 wordList 的长度，C 为列表中单词的长度。哈希表中包含 O(N×C) 个节点，每个节点占用空间 O(C)，因此总的空间复杂度为 O(N×C^2)。
##### Golang实现
    func ladderLength(beginWord string, endWord string, wordList []string) int {
    	wordId := map[string]int{}
    	graph := make([][]int, 0)
    	addWord := func(word string) int {
    		id, ok := wordId[word]
    		if !ok {
    			id = len(wordId)
    			wordId[word] = id
    			graph = append(graph, []int{})
    		}
    		return id
    	}
    	addEdge := func(word string) int {
    		id1 := addWord(word)
    		s := []byte(word)
    		for i, b := range s {
    			s[i] = '*'
    			id2 := addWord(string(s))
    			graph[id1] = append(graph[id1], id2)
    			graph[id2] = append(graph[id2], id1)
    			s[i] = b
    		}
    		return id1
    	}
    
    	for _, word := range wordList {
    		addEdge(word)
    	}
    	beginId := addEdge(beginWord)
    	endId, ok := wordId[endWord]
    	if !ok {
    		return 0
    	}
    
    	const inf int = math.MaxInt64
    	dist := make([]int, len(wordId))
    	for i := range dist {
    		dist[i] = inf
    	}
    	dist[beginId] = 0
    	queue := []int{beginId}
    	for len(queue) > 0 {
    		v := queue[0]
    		queue = queue[1:]
    		if v == endId {
    			return dist[endId]/2 + 1
    		}
    		for _, w := range graph[v] {
    			if dist[w] == inf {
    				dist[w] = dist[v] + 1
    				queue = append(queue, w)
    			}
    		}
    	}
    	
    	return 0
    }
#### 方法二：双向广度优先搜索
##### 复杂度分析
###### 时间复杂度：O(N×C^2)。其中 N 为 wordList 的长度，C 为列表中单词的长度。
###### 空间复杂度：O(N×C^2)。其中 N 为 wordList 的长度，C 为列表中单词的长度。哈希表中包含 O(N×C) 个节点，每个节点占用空间 O(C)，因此总的空间复杂度为 O(N×C^2)。
##### Golang实现
    func ladderLength(beginWord string, endWord string, wordList []string) int {
    	if indexOf(wordList, endWord) == -1 {
    		return 0
    	}
    
    	wordList = append(wordList, beginWord)
    	// 从两端 BFS 遍历要用的队列
    	queue1 := []string{beginWord}
    	queue2 := []string{endWord}
    	// 两端已经遍历过的节点
    	visited1 := map[string]bool{beginWord: true}
    	visited2 := map[string]bool{endWord: true}
    
    	count := 0
    	allWordSet := map[string]bool{}
    	for _, word := range wordList {
    		allWordSet[word] = true
    	}
    
    	for len(queue1) > 0 && len(queue2) > 0 {
    		count++
    
    		if len(queue1) > len(queue2) {
    			queue1, queue2 = queue2, queue1
    			visited1, visited2 = visited2, visited1
    		}
    
    		for i := len(queue1); i > 0; i-- {
    			s := queue1[0]
    			queue1 = queue1[1:]
    			chars := []byte(s)
    
    			for j := 0; j < len(s); j++ {
    				// 保存第j位的原始字符
    				c0 := chars[j]
                    
    				for c := byte('a'); c <= 'z'; c++ {
    					chars[j] = c
    					newString := string(chars)
    					// 已经访问过了，跳过
    					if _, ok := visited1[newString]; ok {
    						continue
    					}
    					// 两端遍历相遇，结束遍历，返回 count
    					if _, ok := visited2[newString]; ok {
    						return count + 1
    					}
    					// 如果单词在列表中存在，将其添加到队列，并标记为已访问
    					if _, ok := allWordSet[newString]; ok {
    						queue1 = append(queue1, newString)
    						visited1[newString] = true
    					}
    				}
    				// 恢复第j位的原始字符
    				chars[j] = c0
    			}
    		}
    	}
    
    	return 0
    }
    
    func indexOf(strList []string, str string) int {
    	for i, strv := range strList {
    		if strv == str {
    			return i
    		}
    	}
    
    	return -1
    }
***
#### 题目
##### 433. 最小基因变化
#### 地址
##### https://leetcode-cn.com/problems/minimum-genetic-mutation/
#### 方法一：单向广度优先搜索
##### 复杂度分析
###### 时间复杂度：O()。
###### 空间复杂度：O()。
##### Golang实现
<pre name="code" class="go">
func minMutation(start string, end string, bank []string) int {
	// 将bank存储到集合中
	geneticSet := map[string]bool{}
	for _, genetic := range bank {
		geneticSet[genetic] = true
	}

	// 检查end是否在集合中
	if _, ok := geneticSet[end]; !ok {
		return -1
	}

	dict := []byte{'A', 'C', 'G', 'T'}
	queue := []string{start}
	delete(geneticSet, start)

	step := 0
	for len(queue) > 0 {
		step++
		for i := len(queue); i > 0; i-- {
			// 出队
			genetic := []byte(queue[0])
			queue = queue[1:]

			for j := 0; j < len(genetic); j++ {
				oldChar := genetic[j] // 临时保存
				for _, dv := range dict {
					genetic[j] = dv // 修改为dict中的字符
					newGenetic := string(genetic)
					if newGenetic == end {
						// 如果和最后一个元素匹配，直接返回
						return step
					}
					if _, ok := geneticSet[newGenetic]; ok {
						// 合法的基因串
						queue = append(queue, newGenetic)
						delete(geneticSet, newGenetic)
					}
				}
				// 还原
				genetic[j] = oldChar
			}
		}
	}

	return -1
}
</pre>
***