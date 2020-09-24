package lru_cache

type LRUCache struct {
	cache    map[int]int
	capacity int
	mark     []int
	count    int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]int, capacity),
		capacity: capacity,
		mark:     make([]int, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.ChangePriority(key)
		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.count == 0 {
		this.count++
		this.mark[this.capacity-1] = key
		this.cache[key] = value
	} else if this.count < this.capacity {
		this.cache[key] = value
		copy(this.mark[this.capacity-this.count-1:this.capacity-1], this.mark[this.capacity-this.count:this.capacity])
		this.mark[this.capacity-1] = key
		this.count++
	} else if this.count == this.capacity {
		if _, ok := this.cache[key]; ok {
			this.cache[key] = value
			this.ChangePriority(key)

		} else {
			keyToDeleted := this.mark[0]
			delete(this.cache, keyToDeleted)
			copy(this.mark[0:], this.mark[1:])
			this.mark[this.capacity-1] = key
			this.cache[key] = value
		}
	}
}

func (this *LRUCache) ChangePriority(key int) {
	for index, markKey := range this.mark {
		if markKey == key {
			copy(this.mark[index:], this.mark[index+1:])
			this.mark[this.capacity-1] = key
			return
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
