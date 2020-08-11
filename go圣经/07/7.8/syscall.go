package main

import (
	"fmt"
)

type Errno uintptr // operating system error code

var errors = [...]string{
	1:   "operation not permitted",   // EPERM
	2:   "no such file or directory", // ENOENT
	3:   "no such process",           // ESRCH
	// ...
}

func (e Errno) Error() string {
	if 0 <= int(e ) && int(e) < len(errors) {
		return errors[e]
	}
	return fmt.Sprintf("errno %d",e)
}

func main() {
	//Errno是一个系统调用错误的高效表示方式，它通过一个有限的集合进行描述，并且它满足标准的错误接口。我们会在第7.11节了解到其它满足这个接口的类型。
	var err error = Errno(4)
	fmt.Println(err.Error()) // "no such file or directory"
	fmt.Println(err)         // "no such file or directory"

	var err2 Errno = 2
	fmt.Println(err2.Error())
}

