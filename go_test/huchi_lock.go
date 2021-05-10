package main

import (
	"fmt"
	"sync"
)

/*
互斥锁

临界区: 当程序同时运行时， 多个Goroutines不能同时访问和修改共享资源代码部分。修改共享资源的这部分代码叫做临界区。


竞争条件解决：
互斥锁 和 channel

使用条件：
一般来说，当 Goroutine 需要彼此通信时使用 channel ，
当需要只有一个 Goroutine 应该访问代码的临界区时使用互斥锁。

 */


var x = 0


// 加锁

//func increment(w *sync.WaitGroup, m *sync.Mutex){
//	m.Lock()
//	x = x + 1
//	m.Unlock()
//	fmt.Println("The value of x: ", x)
//	w.Done()
//}
//
//
//func main()  {
//	var w sync.WaitGroup
//	var m sync.Mutex
//	for i := 0; i < 200; i++ {
//		w.Add(1)
//		go increment(&w, &m)
//	}
//	w.Wait()
//	fmt.Println("finnal value of x: ", x)
//}


// 使用 channel

func increment(w *sync.WaitGroup, ch chan bool){
	ch <- true
	x = x + 1
	<- ch
	w.Done()
}

func main(){
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 200; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("finnal value of x: ", x)
}