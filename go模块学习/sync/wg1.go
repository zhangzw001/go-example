package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go fun1(&wg)
	go fun2(&wg)
	fmt.Println("main 等待子程序执行完毕, 等待中...")
	wg.Wait()
	fmt.Println("子程序全部执行完毕")
}
//
func fun1(wg *sync.WaitGroup) {
	for i := 0 ; i < 5 ; i ++ {
		log.Printf("func1 : %d\n",i)
		time.Sleep(time.Duration(i) * time.Second)
	}
	wg.Done()
}
func fun2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0 ; i < 2 ; i ++ {
		log.Printf("func2 : %d\n",i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}