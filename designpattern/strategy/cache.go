package main

type algorithmType int

const (
	fifoTyp algorithmType = 0 + iota
	lfuTyp
	lruTyp
)

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(algoTyp algorithmType) *Cache {

	var e EvictionAlgo

	switch algoTyp {
	case fifoTyp:
		e = &Fifo{}
	case lfuTyp:
		e = &Lfu{}
	case lruTyp:
		e = &Lru{}
	default:
		panic("algorithm type error.")
	}

	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(algoTyp algorithmType) {

	switch algoTyp {
	case fifoTyp:
		c.evictionAlgo = &Fifo{}
	case lfuTyp:
		c.evictionAlgo = &Lfu{}
	case lruTyp:
		c.evictionAlgo = &Lru{}
	default:
		panic("algorithm type error.")
	}
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

// func (c *Cache) get(key string) {
// 	delete(c.storage, key)
// }

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
