package cache

type Node struct {
		Key, Value string
		Prev, Next *Node
}