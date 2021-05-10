package main

import (
	"fmt"
	"time"
)

//阻塞型管道channel
// 只要没有goroutine 去读取/写入 管道数据， 就发生阻塞

func product(ch chan int){
	for i := 0; i < 10 ; i++{
		ch<- i
	}
	close(ch)  //判断管道状态  关闭为false
}

//func main()  {
//	cha := make(chan int)
//	go product(cha)
//	// 方法1
//	//for {
//	//	v, ok := <- cha
//	//	if ok == false{
//	//		break
//	//	}
//	//	fmt.Println("the result is :", v, "status: ", ok)
//	//}
//	// 方法二  for range
//	for v := range cha{
//		fmt.Println("the result is :", v)
//	}
//}


//非阻塞型管道 channel

// 为管道设置一个缓冲区，也就是设置一个大小
// make(chan int, 2)  为管道设置一个大小为2的缓冲区
// 发生阻塞的情况：
//1. 缓冲区已满，无法写入，阻塞
//2. 管道数据为空，无法读取, 阻塞

func product1(ch chan int){
	for i := 0; i < 5; i++ {
		ch<- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main()  {
	ch := make(chan int, 2)
	go product1(ch)
	time.Sleep(2*time.Second)
	for i := range ch{
		fmt.Println("read value: ", i)
		time.Sleep(2*time.Second)
	}
}