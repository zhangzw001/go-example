package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置一个超时时间
	timeout := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		timeout <- struct{}{}
	}()

	// 设置一个脚本, 如果运行超过timeout时间,则中断(这里模拟程序执行时间)
	ch := make(chan struct{})
	go func() {
		fmt.Println("ch 的子进程开始...")
		rand.Seed(time.Now().Unix())
		r := rand.Intn(10)
		fmt.Printf("正在延迟ch  %ds\n", r)
		time.Sleep(time.Duration(r) * time.Second)
		fmt.Println("ch 的子进程结束...")
		ch <- struct{}{}
	}()

	// 设置一个脚本, 如果运行超过timeout时间,则中断(这里模拟程序执行时间)
	ch2 := make(chan struct{})
	go func() {
		fmt.Println("ch2 的子进程开始...")
		rand.Seed(time.Now().Unix())
		r := rand.Intn(10)
		fmt.Printf("正在延迟ch2  %ds\n", r)
		time.Sleep(time.Duration(r) * time.Second)
		fmt.Println("ch2 的子进程结束...")
		ch2 <- struct{}{}
	}()

	for {
		select {
		case <-ch:
			if _, ok := <-ch2; ok {
				fmt.Println("任务正常结束")
				close(ch)
				return
			}
		case <-time.After(3 * time.Second):
			fmt.Println("time.after")
		case <-timeout:
			fmt.Println("timeout!")
			close(timeout)
			return
		}
	}
}
