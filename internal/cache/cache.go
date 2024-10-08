package cache

import (
	"fmt"
	"main/internal/db"
	"main/internal/types"
	"sync"
	"unsafe"

	"github.com/rs/zerolog/log"
)

var megaByte int = 1024 * 1024

type Cache struct {
	cachedData map[string]types.Order
	m          sync.RWMutex
}

func (c *Cache) SaveData(data types.Order) {
	if c.GetSize() > 200*megaByte {
		clear(c.cachedData)
	}
	c.m.RLock()
	defer c.m.RUnlock()
	c.cachedData[data.OrderUid] = data
}

func (c *Cache) GetOrderById(uid string) (types.Order, error) {
	c.m.RLock()
	defer c.m.RUnlock()
	if order, found := c.cachedData[uid]; found {
		return order, nil
	}
	return types.Order{}, fmt.Errorf("cant find order with id %s in cache", uid)
}

func (c *Cache) GetSize() int {
	return int(unsafe.Sizeof(c.cachedData))
}

func NewCache() *Cache {
	return &Cache{make(map[string]types.Order), sync.RWMutex{}}
}

func (c *Cache) RestoreDataFromDB(d db.DataBase) {
	var data []types.Order = d.GetAllData()
	for _, v := range data {
		c.SaveData(v)
	}
	log.Info().Msg(fmt.Sprintf("Restored %d orders to cache from DB", len(c.cachedData)))
}
