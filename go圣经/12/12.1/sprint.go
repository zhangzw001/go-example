package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case stringer :
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
		///...
		case bool :
			if x {
				return "true"
			}
			return "false"
	default:
		return "???"
	}
}
//但是我们如何处理其它类似[]float64、map[string][]string等类型呢？我们当然可以添加更多的测试分支，但是这些组合类型的数目基本是无穷的。还有如何处理类似url.Values这样的具名类型呢？即使类型分支可以识别出底层的基础类型是map[string][]string，但是它并不匹配url.Values类型，因为它们是两种不同的类型，而且switch类型分支也不可能包含每个类似url.Values的类型，这会导致对这些库的依赖。
//
//没有办法来检查未知类型的表示方式，我们被卡住了。这就是我们为何需要反射的原因。

func main() {
	fmt.Println(reflect.TypeOf(Sprint("1")))
	fmt.Println(reflect.TypeOf(Sprint(1)))
}