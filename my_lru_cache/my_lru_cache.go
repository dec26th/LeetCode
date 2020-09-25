package my_lru_cache

type Node struct {
	Val		int
	Key		int
	Pre		*Node
	Next	*Node
}

type LRUCache struct {
	Num				int
	Head			*Node
	Tail			*Node
	Cache			map[int]int
	MaxCapacity		int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		Num:         0,
		Head:        &Node{},
		Tail:		 nil, // 越靠近尾部表示越在最近被使用
		Cache: 		 make(map[int]int, capacity),
		MaxCapacity: capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	if result, ok := this.Cache[key]; ok {
		this.ModifyNode(key)
		return result
	} else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if this.Num == 0 {
		this.Head.Next = &Node{
			Val:  value,
			Key:  key,
			Pre:  this.Head,
			Next: nil,
		}
		this.Cache[key] = value
		this.Num++
		return
	} else if _, ok := this.Cache[key]; ok {
		this.Cache[key] = value
		this.ModifyNode(key)
		return
	} else {
		if this.Num < this.MaxCapacity {
			this.AddNode(key, value)
			return
		} else {
			delete(this.Cache, this.Head.Next.Key)
			this.Head.Next = this.Head.Next.Next
			this.AddNode(key, value)
		}
	}
}

func (this *LRUCache) ModifyNode(key int) {
	cur := this.Head.Next
	for cur != nil {
		if cur.Key == key {
			this.ModifyPriority(cur)
			return
		}
		cur = cur.Next
	}
}

func (this *LRUCache) ModifyPriority(target *Node) {
	target.Next.Pre = target.Pre
	target.Pre.Next = target.Next

	this.Tail.Next = target
	target.Next = nil

	this.Tail = this.Tail.Next
}

func (this *LRUCache) AddNode(key, value int) {
	this.Cache[key] = value
	this.Tail.Next = &Node{
		Val:  value,
		Key:  key,
		Pre:  this.Tail,
		Next: nil,
	}
	this.Tail = this.Tail.Next
	this.Num++
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

