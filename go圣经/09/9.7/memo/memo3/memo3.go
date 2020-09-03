package memo3

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

/*
这些修改使性能再次得到了提升，但有一些URL被获取了两次。这种情况在两个以上的goroutine同一时刻调用Get来请求同样的URL时会发生。多个goroutine一起查询cache，发现没有值，然后一起调用f这个慢不拉叽的函数。在得到结果后，也都会去更新map。其中一个获得的结果会覆盖掉另一个的结果。

*/
func ( memo *Memo) Get(key string ) ( interface{}, error ) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()

	if !ok {
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}

	return res.value, res.err
}