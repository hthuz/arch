

# Trie

前缀树

根节点不包含字符，除根节点外每一个节点都只包含一个字符。
从根节点到某一节点，路径上经过的字符连接起来，为该节点对应的字符串。
每个节点的所有子节点包含的字符都不相同。

假如给定三个单词，
apple
app
ape
那么他们就会组成trie

```
(root)
 └── a
      └── p
           ├── p
           │     └── l
           │           └── e
           └── e
           
```





## 实现

最基本的操作，插入一个字符串，判断字符串是否存在（这个哈希也可以做到），判断某个prefix出现的次数

```go
package main

type TrieNode struct {
	children [26]*TrieNode
	pass     int // 经过该节点次数
	end      int // 结束次数, 相当于这个单词出现的次数
}


// => 每个节点26个children费空间，也可以
// type TrieNode struct {
//  chidren map[rune]*TrieNode
// }

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, w := range word {
		c := w - 'a'
		node.pass += 1

		if node.children[c] == nil {
			node.children[c] = &TrieNode{}
		}
		node = node.children[c]
	}

	node.end += 1

}

func (t *Trie) ContainString(word string) bool {
	node := t.root
	for _, w := range word {
		c := w - 'a'
		if node.children[c] == nil {
			return false
		}
		node = node.children[c]
	}
	return node.end > 0

}

// 某个prefix出现的次数
func (t *Trie) PreFixCount(prefix string) int {
	node := t.root
	for _, w := range prefix {
		c := w - 'a'
		if node.children[c] == nil {
			return 0
		}
		node = node.children[c]
	}
	return node.pass

}

```



## 自动补全

```go
// 比如基于trie实现的自动补全
// => 给定一个prefix， 显示出基于此prefix的所有单词
// 实际可能每个Trie节点都会记录top k结果，方便快速返回
// 以及incremental search， 当用户输入是连续的，就没有必要每次都从root开始找
func (t *Trie) AutoCompletion(prefix string) []string {

	node := t.root

	for _, w := range prefix {
		c := w - 'a'
		if node.children[c] == nil {
			return nil
		}
		node = node.children[c]
	}

	res := make([]string, 0)
	var dfs func(node *TrieNode, word string)

	dfs = func(node *TrieNode, word string) {
		if node.end > 0 {
			res = append(res, word)
		}

		for i, next := range node.children {
			if next != nil {
				w := string(rune('a' + i))
				dfs(next, word+w)
			}
		}
	}
	dfs(node, prefix)

	return res

}

```





## 示例

leetcode 648 单词替换
