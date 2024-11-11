package cache

// LRUCache struct manages the LRU cache with capacity control and node operations.
type LRUCache struct {
    Capacity int
    Cache    map[string]*Node
    Head, Tail *Node
}

// NewLRUCache initializes a new LRU cache with given capacity.
func NewLRUCache(capacity int) *LRUCache {
    cache := &LRUCache{
        Capacity: capacity,
        Cache:    make(map[string]*Node),
    }
    cache.Head = &Node{}
    cache.Tail = &Node{}
    cache.Head.Next = cache.Tail
    cache.Tail.Prev = cache.Head
    return cache
}

// Get retrieves a value by key, returning it and updating access order.
func (c *LRUCache) Get(key string) (string, bool) {
    if node, found := c.Cache[key]; found {
        c.moveToFront(node)
        return node.Value, true
    }
    return "", false
}

// Set inserts or updates a key-value pair in the cache.
func (c *LRUCache) Set(key, value string) {
    if node, found := c.Cache[key]; found {
        node.Value = value
        c.moveToFront(node)
    } else {
        if len(c.Cache) == c.Capacity {
            c.removeOldest()
        }
        newNode := &Node{Key: key, Value: value}
        c.Cache[key] = newNode
        c.addToFront(newNode)
    }
}

// moveToFront moves a node to the front (most recently used).
func (c *LRUCache) moveToFront(node *Node) {
    c.remove(node)
    c.addToFront(node)
}

// remove removes a node from the linked list.
func (c *LRUCache) remove(node *Node) {
    prev := node.Prev
    next := node.Next
    prev.Next = next
    next.Prev = prev
}

// addToFront adds a node right after head (most recently used).
func (c *LRUCache) addToFront(node *Node) {
    node.Prev = c.Head
    node.Next = c.Head.Next
    c.Head.Next.Prev = node
    c.Head.Next = node
}

// removeOldest removes the least recently used node.
func (c *LRUCache) removeOldest() {
    oldest := c.Tail.Prev
    c.remove(oldest)
    delete(c.Cache, oldest.Key)
}
