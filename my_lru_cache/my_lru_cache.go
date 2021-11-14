package my_lru_cache

type Node struct {
	Val  int
	Key  int
	Pre  *Node
	Next *Node
}

type LRUCache struct {
	Num         int
	Head        *Node
	Tail        *Node
	Cache       map[int]*Node
	MaxCapacity int
}

func Constructor(capacity int) LRUCache {
	LRUCache := LRUCache{
		Num:         0,
		Head:        &Node{},
		Tail:        &Node{}, // 越靠近尾部表示越在最近被使用
		Cache:       make(map[int]*Node, capacity),
		MaxCapacity: capacity,
	}
	LRUCache.Head.Next = LRUCache.Tail
	LRUCache.Tail.Pre = LRUCache.Head
	return LRUCache
}

func (this *LRUCache) Get(key int) int {
	if result, ok := this.Cache[key]; ok {
		this.ModifyPriority(result)
		return result.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; !ok {
		if this.Num != this.MaxCapacity {
			this.AddNode(key, value)
			return
		} else {
			this.Num--
			delete(this.Cache, this.Tail.Pre.Key)
			this.RemoveNode(this.Tail.Pre)
			this.AddNode(key, value)
			return
		}
	} else {
		node.Val = value
		this.ModifyPriority(node)
	}
}

func (this *LRUCache) ModifyPriority(result *Node) {
	this.RemoveNode(result)

	result.Pre = this.Head
	this.Head.Next.Pre = result
	result.Next = this.Head.Next
	this.Head.Next = result
}

func (this *LRUCache) AddNode(key int, value int) {
	node := &Node{
		Val: value,
		Key: key,
	}

	this.Head.Next.Pre = node
	node.Next = this.Head.Next
	this.Head.Next = node
	node.Pre = this.Head

	this.Num++
	this.Cache[key] = node
}

func (this *LRUCache) RemoveNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
