package memo4

import (
	"fmt"
	"sync"
)

type entry struct {
	res 	result
	ready 	chan struct{}
}

type Memo struct {
	f 		Func
	mu 		sync.Mutex
	cache 	map[string]*entry
}

type Func func(key string ) ( interface{}, error )

type result struct {
	value 	interface{}
	err 	error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

/*
这些修改使性能再次得到了提升，但有一些URL被获取了两次。这种情况在两个以上的goroutine同一时刻调用Get来请求同样的URL时会发生。多个goroutine一起查询cache，发现没有值，然后一起调用f这个慢不拉叽的函数。在得到结果后，也都会去更新map。其中一个获得的结果会覆盖掉另一个的结果。

*/
func ( memo *Memo) Get(key string ) (value interface{}, err error ) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil  {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready : make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		fmt.Println("不存在的key")
		close(e.ready)
	}else {
		memo.mu.Unlock()
		fmt.Println("存在的key")
		<- e.ready
	}

	return e.res.value, e.res.err
}