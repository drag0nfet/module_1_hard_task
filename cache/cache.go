package cache

import "sync"

// Interface - реализуйте этот интерфейс
type Interface interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// Не меняйте названия структуры и название метода создания экземпляра Cache, иначе не будут проходить тесты

type Cache struct {
	storage map[string]string
	mutex   sync.RWMutex
}

// NewCache создаёт и возвращает новый экземпляр Cache.
func NewCache() Interface {
	return &Cache{
		storage: make(map[string]string),
		mutex:   sync.RWMutex{},
	}
}

func (c *Cache) Set(k, v string) {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	c.storage[k] = v
}

func (c *Cache) Get(k string) (v string, ok bool) {
	defer c.mutex.RUnlock()
	c.mutex.RLock()
	v, ok = c.storage[k]
	return v, ok
}
