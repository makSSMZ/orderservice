package internal

import "sync"

// структура кеша
type Cache struct {
	mu     sync.RWMutex // просто синхронизация доступа, к примеру, либо блочим запись, либо отпускаем
	orders map[string]Order
}

func NewCache() *Cache {
	return &Cache{
		orders: make(map[string]Order),
	}
}

func (c *Cache) Get(id string) (Order, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	order, ok := c.orders[id]
	return order, ok
}

func (c *Cache) Set(order Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.orders[order.OrderUID] = order
}

func (c *Cache) Load(orders []Order) {
	for _, o := range orders {
		c.Set(o)
	}
}
