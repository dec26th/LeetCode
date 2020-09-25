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
	MaxCapacity		int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		Num:         0,
		Head:        &Node{},
		Tail:		 nil, // 越靠近尾部表示越在最近被使用
		MaxCapacity: capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	cur := this.Head.Next
	for cur != nil {
		if cur.Key == key {

			cur.Pre.Next = cur.Next
			cur.Next.Pre = cur.Pre

			this.Tail.Next = cur
			this.Tail.Next.Pre = this.Tail
			cur.Next = nil
			this.Tail = this.Tail.Next
			return cur.Val
		}

		cur = cur.Next
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if this.Num != this.MaxCapacity {

		if this.Num == 0 {
			this.Tail = &Node {
				Val:  value,
				Key:  key,
				Pre:  this.Head,
				Next: nil,
			}
			this.Head.Next = this.Tail
			this.Num++
			return
		}

		this.Tail.Next = &Node{
			Val:  value,
			Key:  key,
			Pre:  this.Tail,
			Next: nil,
		}
		this.Tail = this.Tail.Next
		this.Num++
		return
	} else { // 超过容量限制
		cur := this.Head.Next  // 头节点
		for cur != nil {
			if cur.Key == key {
				cur.Val = value

				cur.Next.Pre = cur.Pre
				cur.Pre.Next = cur.Next

				this.Tail.Next = cur
				cur.Next = nil
				this.Tail.Next.Pre = this.Tail
				this.Tail = this.Tail.Next

				return
			}
			cur = cur.Next
		}

		this.Head.Next = this.Head.Next.Next  //head是无用节点， head.next才是实际的头节点
		this.Tail.Next = &Node{
			Val:  value,
			Key:  key,
			Pre:  this.Tail,
			Next: nil,
		}
		this.Tail = this.Tail.Next
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

