package hashTable

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type IHashTable interface {
	Put(key interface{}, value interface{}) error
	Get(key interface{}) *Entry
	Remove(key interface{}) error
	Hash(key interface{}) int
	Show()
}

type Options struct {
	// hashtable容量，设置默认桶容量
	Capacity   uint
	// 负载因子 0<=x<=1
	LoadFactor float64
	Debug      bool
}

type Entry struct {
	key   interface{}
	value interface{}
	next  *Entry
}

type HashTable struct {
	// hash桶
	// Hashtable保存key-value的数组。
	// Hashtable是采用拉链法实现的，每一个Entry本质上是一个单向链表
	table []*Entry

	// hashtable的设定容量
	capacity int

	// Hashtable中元素的实际数量
	count int

	// 阈值，用于判断是否需要调整Hashtable的容量（threshold = 容量*加载因子）
	threshold int

	// 加载因子, 超过此比例就自动扩容
	loadFactor float64

	// debug
	debug bool

	lock sync.Mutex
}

var _ IHashTable = &HashTable{}

func NewHashTable(o *Options) *HashTable {
	var cap = o.Capacity
	if cap == 0 {
		cap = 1
	}
	var ht = new(HashTable)
	ht.loadFactor = o.LoadFactor
	ht.capacity = int(o.Capacity)
	ht.threshold = int(o.LoadFactor * float64(cap))
	ht.table = make([]*Entry, cap)
	ht.debug = o.Debug
	return ht
}

func (h *HashTable) Put(key interface{}, value interface{}) error {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.insert(key, value)

	h.count++

	// 检查是否达到了设定的负载值
	if int(h.count) >= h.threshold {
		h.reHash()
	}

	return nil
}

func (h *HashTable) insert(key interface{}, value interface{}) error {
	// 找出key所在的table
	hash := h.Hash(key)

	var entry = h.table[hash]
	if entry == nil {
		h.table[hash] = &Entry{key: key, value: value}
	} else if entry.key == key {
		entry.value = value
	} else {
		for entry.next != nil {
			// 如果已存在该key，则修改
			if entry.next.key == key {
				entry.next.value = value
				return nil
			}
			entry.next = entry.next.next
		}

		entry.next = &Entry{key: key, value: value}
	}

	return nil
}

func (h *HashTable) reHash() error {
	var oldTable = h.table
	var oldCap = h.capacity
	// 新的容量为
	var newCap = oldCap*2 + 1
	// 设置新的门槛
	h.threshold = int(float64(newCap) * h.loadFactor)
	// 设置新的容量
	h.capacity = newCap
	// 设置新的entry
	var newEntry = make([]*Entry, newCap)
	h.table = newEntry

	// 记录log
	if h.debug {
		log.Printf("数量到达: %v(%v*%v)，开始扩容: %v -> %v \n",
			h.count, oldCap, h.loadFactor, oldCap, newCap)
	}

	// 将旧的entry复制到新entry
	for i := oldCap - 1; i >= 0; i-- {
		current := oldTable[i]
		for current != nil {
			//h.table[h.Hash(current.key)] = &Entry{
			//	key:   current.key,
			//	value: current.value,
			//}
			h.insert(current.key, current.value)
			current = current.next
		}
	}
	return nil
}

func (h *HashTable) Get(key interface{}) *Entry {
	// 获取key的hash所在table
	hash := h.Hash(key)
	// 获取对应的entry
	entry := h.table[hash]

	if entry == nil {
		return nil
	}

	for entry.key != key && entry.next != nil {
		entry = entry.next
	}

	if entry.key == key {
		return entry
	}

	return nil
}

func (h *HashTable) Remove(key interface{}) error {
	// 获取key的hash所在table
	hash := h.Hash(key)
	// 获取对应的entry
	var prev = h.table[hash]

	if prev == nil {
		return errors.New("the given key is not exists")
	}

	if prev.key == key {
		h.table[hash] = prev.next
		//prev = nil
		return nil
	}

	for prev.key != key && prev.next != nil {
		if prev.next.key == key {
			prev.next = prev.next.next
			//prev = nil
			return nil
		}
		prev = prev.next
	}
	return errors.New("the given key is not exists")
}

func (h *HashTable) Hash(key interface{}) int {
	return HashCode(key) % h.capacity
}

func (h *HashTable) Show() {
	for k, item := range h.table {
		fmt.Printf("%v: ", k)
		for item != nil {
			fmt.Printf("%v=>%v, ", item.key, item.value)
			item = item.next
		}
		fmt.Println("")
	}
}

func HashCode(key interface{}) int {
	keyStr := fmt.Sprintf("%s", key)
	keyLen := len(keyStr)

	var h = 0

	if keyLen > 0 {
		for i := 0; i < keyLen; i++ {
			h = h<<5 - h + int(keyStr[i])
		}
	}

	return h
}
