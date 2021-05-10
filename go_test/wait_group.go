package main

import (
	"fmt"
	"sync"
	"time"
)

// waitGroup:等待所有的Goroutines完成执行
// 工作原理是：计数器


func process(i int, wg *sync.WaitGroup){
	fmt.Println("start goroutine...", i)
	time.Sleep(2*time.Second)
	fmt.Println("=============")
	fmt.Println("finished goroutine...")
	wg.Done()  // 递减计数器: wg里的数字会递减
}



func main()  {
	num := 3
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)  // 递增计数器: Add 传进来的int 会递增
		fmt.Println("i:",i)
		//传递wg地址进来很重要，如果不传递wg地址，则每个goroutine都会生成自己的waitGroup副本，当他们执行完毕时，main不会收到通知；
		go process(i, &wg)
		fmt.Println("ii:",i)

	}
	wg.Wait()  // main函数 阻塞调用它的Goroutine,直到计数器变为0
	fmt.Println("All go routines is finished...")
}