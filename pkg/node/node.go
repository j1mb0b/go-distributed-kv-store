package node

import "sync"

type Node struct {
	ID string
	data map[string]string
	mu sync.RWMutex
}

func NewNode(id string) *Node {
	return &Node{
		ID: id,
		data: make(map[string]string),
	}
}

func (n *Node) Put(key, value string) {
    n.mu.Lock()
    defer n.mu.Unlock()
    n.data[key] = value
}

func (n *Node) Get(key string) (string, bool) {
    n.mu.RLock()
    defer n.mu.RUnlock()
    value, ok := n.data[key]
    return value, ok
}
