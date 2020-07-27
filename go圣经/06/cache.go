package main

import (
	"fmt"
	"sync"
)

// 1 普通方式
var (
	mu 	sync.Mutex
	mapping = make(map[string]string )
)

func lookup(key string ) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

// 2 类方式
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{ mapping: make(map[string]string )}

//我们给新的变量起了一个更具表达性的名字：cache。因为sync.Mutex字段也被嵌入到了这个struct里，其Lock和Unlock方法也就都被引入到了这个匿名结构中了，这让我们能够以一个简单明了的语法来对其进行加锁解锁操作。
func lookup2(key string ) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

// 3 方法
type Cache struct {
	mu sync.Mutex
	mapping map[string]string
}

func (c Cache) lookup3(key string ) string  {
	c.mu.Lock()
	v := c.mapping[key]
	c.mu.Unlock()
	return v
}


func main() {
	cache.mapping["aa"] = "111"
	fmt.Println(lookup2("aa"))
	var c Cache
	c.mapping = make(map[string]string)
	c.mapping["bb"]  = "222"
	fmt.Println(c.lookup3("bb"))
}
