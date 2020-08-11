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
	var err error = Errno(2)
	fmt.Println(err.Error()) // "no such file or directory"
	fmt.Println(err)         // "no such file or directory"

	var err2 Errno = 1
	fmt.Println(err2.Error())
}

