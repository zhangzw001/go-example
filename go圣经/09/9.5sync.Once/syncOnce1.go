package main

import (
	"image"
	"os"
	"sync"
)

var icons map[string]image.Image

func loadIcon(s string) image.Image  {
	f , _ := os.Open(s)
	defer f.Close()
	img , _ , _ := image.Decode(f)
	return img
}

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

// NOTE1 : not concurrency-safe!
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons() // one-time initialization
	}
	return icons[name]
}

// NOTE2: Concurrency-safe.
var mu sync.Mutex

func Icon2(name string ) image.Image {
	mu.Lock()
	defer mu.Unlock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

// NOTE3: Concurrency-safe.
// 然而使用互斥访问icons的代价就是没有办法对该变量进行并发访问，即使变量已经被初始化完毕且再也不会进行变动。这里我们可以引入一个允许多读的锁：
var murw sync.RWMutex
func Icon3(name string ) image.Image {
	if icons != nil {
		murw.RLock()
		defer murw.RUnlock()
		return icons[name]
	}

	murw.Lock()
	defer murw.Unlock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

// NOTE4: Concurrency-safe.
// sync.Once。概念上来讲，一次性的初始化需要一个互斥量mutex和一个boolean变量来记录初始化是不是已经完成了；互斥量用来保护boolean变量和客户端数据结构。Do这个唯一的方法需要接收初始化函数作为其参数
var loadIconsOnce sync.Once
func Icon4(name string ) image.Image {
	loadIconsOnce.Do(loadIcons)
	return  icons[name]
}