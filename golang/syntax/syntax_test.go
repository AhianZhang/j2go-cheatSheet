package main

import (
	"fmt"
	"testing"
	"time"
)

// new function
func TestNewFunc(t *testing.T) {
	username1 := new(string)
	username2 := new(string)
	t.Log(username1)
	t.Log(username2)
}

// time
func TestTime(t *testing.T) {
	now := time.Now() // 当前时间 默认格式：2006-01-02 15:04:05.999999999 -0700 MST
	fmt.Println(now)
	format := now.Format("2006-01-02 15:04:05") // 格式化 年月日时分秒
	fmt.Println(format)
}

type Trie struct {
	IsEnd    bool
	Children [26]*Trie
}

func (trie *Trie) Insert(word string) {
	node := trie
	for _, ch := range word {
		ch -= 'a'
		if node.Children[ch] == nil {
			node.Children[ch] = &Trie{}
		}
		node = node.Children[ch]
	}
	node.IsEnd = true
	fmt.Println(trie)
}

func TestTrie(t *testing.T) {
	trie := Trie{}
	trie.Insert("asdab")
	fmt.Println("contains word", trie.Search("asda"))
	println("contains word", trie.Search("asdab"))
	fmt.Println("contains prefix", trie.StartWith("asda"))

}

func (trie *Trie) SearchPrefix(word string) *Trie {
	node := trie
	for _, ch := range word {
		ch -= 'a'
		if node.Children[ch] == nil {
			return nil
		}
		node = node.Children[ch]

	}
	return node
}

func (trie *Trie) Search(word string) bool {
	node := trie.SearchPrefix(word)
	if node != nil && node.IsEnd == true {
		return true
	}
	return false
}

func (trie *Trie) StartWith(word string) bool {
	return trie.SearchPrefix(word) != nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recur(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
		fmt.Println("root is nil")
	}
	fmt.Println(root.Val)
	root.Left = recur(root.Left)
	root.Right = recur(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}
func TestTreeNode(t *testing.T) {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	recur(node)
}
