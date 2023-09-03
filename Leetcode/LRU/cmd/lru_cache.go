package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {
	fmt.Println("vim-go")
	logger := zap.NewExample()
	logger.Info("start")
	cache := NewCache(5)
	cache.Put("A", "A")
	cache.Put("B", "B")
	cache.Put("C", "C")
	cache.Put("D", "D")
	cache.Put("E", "E")
	cache.Put("A", "A")
	cache.Put("F", "F")
	cache.Put("B", "B")
	cache.Put("T", "T")
	cache.Put("A", "A")
	// logger.Info("cache", zap.Any("cache", cache))
	fmt.Println(len(cache.Hash))
	node := cache.Head
	for node.Next != nil {
		node = node.Next
		fmt.Printf("%s,%s\n", node.Key, node.Val)
	}
}

type Node struct {
	Pre  *Node
	Next *Node
	Val  string
	Key  string 
}

func NewNode(Key string, Val string) *Node {
	return &Node{
		Key: Key,
		Val: Val,
	}
}

type HandmadeCache struct {
	Head     *Node
	Tail     *Node
	Hash     map[string]*Node
	Capacity int
}

func NewCache(capacity int) *HandmadeCache {
	head := NewNode("", "")
	tail := NewNode("", "")
	head.Next = tail
	tail.Pre = head
	return &HandmadeCache{
		Head:     head,
		Tail:     tail,
		Hash:     map[string]*Node{},
		Capacity: capacity,
	}
}

func (cache *HandmadeCache) Put(key string, val string) {
	if _, ok := cache.Hash[key]; ok {
		cache.remove(key)
	}
	if len(cache.Hash) >= cache.Capacity {
		cache.remove(cache.Head.Next.Key)
	}
	cache.insert(NewNode(key, val))
}

func (cache *HandmadeCache) remove(key string) {
	if len(cache.Hash) == 0 {
		return
	}
	if node, ok := cache.Hash[key]; ok {
		delete(cache.Hash, key)
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	}
}

func (cache *HandmadeCache) insert(node *Node) {
	cache.Hash[node.Key] = node
	lastNode := cache.Tail.Pre
	lastNode.Next = node
	node.Pre = lastNode
	node.Next = cache.Tail
	cache.Tail.Pre = node
}

func (cache *HandmadeCache) Get(key string) string {
	if node, ok := cache.Hash[key]; ok {
		cache.remove(key)
		cache.insert(node)
		return node.Val
	}
	return ""
}
