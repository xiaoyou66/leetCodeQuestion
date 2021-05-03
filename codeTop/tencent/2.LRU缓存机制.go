// Package tencent
// @Description
// @Author 小游
// @Date 2021/04/15
package tencent

import "container/list"

type entry struct {
	key, value int
}

type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	lst   *list.List
}

// 构造缓存
func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, map[int]*list.Element{}, list.New()}
}

// 获取值
func (c *LRUCache) Get(key int) int {
	// 获取值，如果不为空，那么就返回-1
	e := c.cache[key]
	if e == nil {
		return -1
	}
	// 把当前值移动到头部
	c.lst.MoveToFront(e) // 刷新缓存使用时间
	return e.Value.(entry).value
}

// 设置值
func (c *LRUCache) Put(key, value int) {
	// 判断缓存中是否存在
	if e := c.cache[key]; e != nil {
		// 存在就更新
		e.Value = entry{key, value}
		c.lst.MoveToFront(e) // 刷新缓存使用时间
		return
	}
	// 不存在我们就放一个新的
	c.cache[key] = c.lst.PushFront(entry{key, value})
	// 哦按的是否超过大小，如果超过，那么就删除
	if len(c.cache) > c.cap {
		delete(c.cache, c.lst.Remove(c.lst.Back()).(entry).key)
	}
}
