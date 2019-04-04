package hashTable

type IHashTable interface {
	Add()
	Get()
	Delete()
	Hash()
	Show()
}

type HashNode struct {
	key   int
	value interface{}
	next  *HashNode
}

type HashTable struct {
	// 默认hash桶容量
	cap int
	// 负载因子,超过此比例就自动扩容,默认 0.75
	rate float64
	// 最大hash桶容量
	maxCap int
	// hash桶
	buckets []HashNode
	// 当前存放节点数
	size int
}

type Options struct {
	Cap int
	Rate float64
	MaxCap int
}

func NewHashTable(options *Options) *HashTable {
	return &HashTable{
		cap:options.Cap,
		rate:options.Rate,
		maxCap:options.MaxCap,
		buckets:make([]HashNode,options.Cap,options.MaxCap),
	}
}

func (h *HashTable) Add(key interface{}, value interface{}) error {

	return nil
}

func (h *HashTable) Hash(key interface{}) int {

	return 0
}
