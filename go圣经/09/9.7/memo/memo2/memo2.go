package memo2

import "sync"

type Memo struct {
	f 		Func
	mu 		sync.Mutex
	cache 	map[string]result
}

type Func func(key string ) ( interface{}, error )

type result struct {
	value 	interface{}
	err 	error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

/*测试依然并发进行，但这回竞争检查器“沉默”了。不幸的是对于Memo的这一点改变使我们完全丧失了并发的性能优点。每次对f的调用期间都会持有锁，Get将本来可以并行运行的I/O操作串行化了。我们本章的目的是完成一个无锁缓存，而不是现在这样的将所有请求串行化的函数的缓存。
*/
func ( memo *Memo) Get(key string ) ( interface{}, error ) {
	memo.mu.Lock()
	defer memo.mu.Unlock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}

	return res.value, res.err
}