package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type KeyValue struct {
	Key   Key
	Value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	kv := KeyValue{Key: key, Value: value}
	item, ok := c.items[key]
	if ok {
		item.Value = kv
		c.queue.MoveToFront(item)
		return true
	}

	if c.queue.Len()+1 > c.capacity {
		backItem := c.queue.Back()
		c.queue.Remove(backItem)
		delete(c.items, backItem.Value.(KeyValue).Key)
	}
	c.items[key] = c.queue.PushFront(kv)
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	c.queue.MoveToFront(item)
	return item.Value.(KeyValue).Value, true
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
